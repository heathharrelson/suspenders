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
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gorilla/mux"

	"k8s.io/apimachinery/pkg/labels"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	appsinformers "k8s.io/client-go/informers/apps/v1"
	"k8s.io/client-go/kubernetes"
	appslisters "k8s.io/client-go/listers/apps/v1"
	"k8s.io/client-go/tools/cache"
	klog "k8s.io/klog/v2"
)

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
	indexTemplate, err := template.New("index.html").ParseFiles("ui/dist/index.html")
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

	router := mux.NewRouter()
	pathPrefix := s.contextPath()

	router.HandleFunc("/healthz", s.handleHealth)

	appRouter := router
	if pathPrefix != "" {
		appRouter = router.PathPrefix(pathPrefix).Subrouter()
	}

	appRouter.HandleFunc("/api/v1/deployments", s.handleListDeployments).Methods("GET")

	fs := http.FileServer(http.Dir("ui/dist"))
	assetPaths := []string{"/css/", "/js/", "/img/"}
	for _, path := range assetPaths {
		appRouter.PathPrefix(path).Handler(http.StripPrefix(pathPrefix, fs))
	}

	appRouter.PathPrefix("/").Handler(http.HandlerFunc(s.handleIndex))

	klog.Info("Starting server on port 8080")
	klog.Info("External URL:", s.externalURL)

	srv := http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
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
	templateData := make(map[string]interface{})
	templateData["contextPath"] = s.contextPath()

	err := s.template.Execute(w, templateData)
	if err != nil {
		klog.Fatal(err.Error())
	}
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ok")
}

func (s *Server) handleListDeployments(w http.ResponseWriter, r *http.Request) {
	deployList, err := s.deploymentLister.List(labels.Everything())
	if err != nil {
		http.Error(w, "Error listing deployments", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(deployList)
}
