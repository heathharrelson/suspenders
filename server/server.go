package server

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"
	"strings"

	humanize "github.com/dustin/go-humanize"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
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
	clientset *kubernetes.Clientset
	template  *template.Template
}

// NewServer creates a new HTTP server
func NewServer(clientset *kubernetes.Clientset) *Server {
	indexTemplate, err := template.New("index.html").Funcs(funcs).ParseFiles("server/templates/index.html")
	if err != nil {
		panic(err)
	}

	return &Server{
		clientset: clientset,
		template:  indexTemplate,
	}
}

// Run starts the HTTP server
func (s *Server) Run() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		deployList, err := s.clientset.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			http.Error(w, "Error listing deployments", http.StatusInternalServerError)
			return
		}

		rows := deploymentRows(deployList)
		err = s.template.Execute(w, rows)
		if err != nil {
			log.Panic(err)
		}
	})

	log.Printf("Serving on port 8080...")
	return http.ListenAndServe(":8080", nil)
}

func deploymentRows(dl *appsv1.DeploymentList) []deploymentRow {
	rows := make([]deploymentRow, len(dl.Items))

	for i, d := range dl.Items {
		lc := latestCondition(d.Status.Conditions)
		rows[i] = deploymentRow{
			Name:               fmt.Sprintf("%s/%s", d.Namespace, d.Name),
			Images:             images(d),
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
