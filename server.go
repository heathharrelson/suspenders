/*
Copyright 2020 Heath Harrelson.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"sort"
	"strings"

	humanize "github.com/dustin/go-humanize"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	appsinformers "k8s.io/client-go/informers/apps/v1"
	"k8s.io/client-go/kubernetes"
	appslisters "k8s.io/client-go/listers/apps/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
)

var funcs = template.FuncMap{
	"relative": func(t metav1.Time) string {
		return humanize.Time(t.Time)
	},
}

type deploymentRow struct {
	Name               string
	Images             string
	DesiredReplicas    int32
	ReadyReplicas      int32
	UpdatedReplicas    int32
	CreationTimestamp  metav1.Time
	LastTransitionTime metav1.Time
	Status             string
}

// Server handles HTTP requests
type Server struct {
	clientset          kubernetes.Interface
	deploymentInformer appsinformers.DeploymentInformer
	deploymentLister   appslisters.DeploymentLister
	deploymentsSynced  cache.InformerSynced
	externalURL        *url.URL
	template           *template.Template
}

// NewServer creates a new HTTP server
func NewServer(clientset kubernetes.Interface, deploymentInformer appsinformers.DeploymentInformer, externalURL *url.URL) *Server {
	indexTemplate, err := template.New("index.html").Funcs(funcs).ParseFiles("templates/index.html")
	if err != nil {
		panic(err)
	}

	return &Server{
		clientset:          clientset,
		deploymentInformer: deploymentInformer,
		deploymentLister:   deploymentInformer.Lister(),
		deploymentsSynced:  deploymentInformer.Informer().HasSynced,
		externalURL:        externalURL,
		template:           indexTemplate,
	}
}

// Run starts the HTTP server
func (s *Server) Run(stopCh <-chan struct{}) error {
	defer utilruntime.HandleCrash()

	klog.Info("Starting the web controller")

	klog.Info("Waiting for informer caches to sync")
	if ok := cache.WaitForCacheSync(stopCh, s.deploymentsSynced); !ok {
		return fmt.Errorf("failed to wait for caches to sync")
	}

	srv := http.Server{Addr: ":8080"}

	assetPath := s.contextPath() + "/static/"
	fs := http.FileServer(http.Dir("ui/static"))
	http.Handle(assetPath, http.StripPrefix(assetPath, fs))

	indexPath := s.contextPath() + "/"
	http.HandleFunc(indexPath, s.handleIndex)

	http.HandleFunc("/healthz", s.handleHealthCheck)

	klog.Info("Starting server on port 8080")
	klog.Info("External URL:", s.externalURL)
	go func() { srv.ListenAndServe() }()

	<-stopCh

	klog.Info("Shutting down HTTP server")
	if err := srv.Shutdown(context.Background()); err != nil {
		return fmt.Errorf("HTTP server Shutdown: %v", err)
	}

	return nil
}

func (s *Server) contextPath() string {
	return strings.TrimRight(s.externalURL.Path, "/")
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	deployList, err := s.deploymentLister.List(labels.Everything())
	if err != nil {
		http.Error(w, "Error listing deployments", http.StatusInternalServerError)
		return
	}

	templateData := make(map[string]interface{})
	templateData["contextPath"] = s.contextPath()
	templateData["deploymentRows"] = deploymentRows(deployList)

	err = s.template.Execute(w, templateData)
	if err != nil {
		klog.Fatal(err.Error())
	}
}

func (s *Server) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ok")
}

func deploymentRows(dl []*appsv1.Deployment) []deploymentRow {
	rows := make([]deploymentRow, len(dl))

	for i, d := range dl {
		lc := latestCondition(d.Status.Conditions)
		rows[i] = deploymentRow{
			Name:               fmt.Sprintf("%s/%s", d.Namespace, d.Name),
			Images:             images(*d),
			DesiredReplicas:    d.Status.Replicas,
			ReadyReplicas:      d.Status.ReadyReplicas,
			UpdatedReplicas:    d.Status.UpdatedReplicas,
			CreationTimestamp:  d.CreationTimestamp,
			LastTransitionTime: lc.LastTransitionTime,
			Status:             fmt.Sprintf("%s:%s", lc.Type, lc.Status),
		}
	}

	sort.Slice(rows, func(i, j int) bool {
		return rows[i].LastTransitionTime.Before(&rows[j].LastTransitionTime)
	})

	return rows
}

func images(d appsv1.Deployment) string {
	var s []string

	for _, c := range d.Spec.Template.Spec.Containers {
		s = append(s, c.Image)
	}

	return strings.Join(s, ", ")
}

func latestCondition(conditions []appsv1.DeploymentCondition) appsv1.DeploymentCondition {
	max := conditions[0]
	for _, condition := range conditions {
		if max.LastTransitionTime.Before(&condition.LastTransitionTime) {
			max = condition
		}
	}

	return max
}
