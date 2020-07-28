// Iniitially derived from https://github.com/kubernetes/client-go/tree/master/examples/out-of-cluster-client-configuration
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	//"k8s.io/apimachinery/pkg/api/errors"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

func main() {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	for {
		deployments, err := clientset.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d deployments in the cluster\n", len(deployments.Items))
		for _, deployment := range deployments.Items {
			fmt.Printf("Name: %s.%s\n", deployment.Namespace, deployment.Name)
			fmt.Printf("Created at: %v\n", deployment.CreationTimestamp)
			fmt.Printf("Desired replicas: %d\n", deployment.Status.Replicas)
			fmt.Printf("Ready replicas: %d\n", deployment.Status.ReadyReplicas)

			lastCond := latestCondition(deployment.Status.Conditions)
			fmt.Printf("Last status: %v:%v at %v\n", lastCond.Type, lastCond.Status, lastCond.LastTransitionTime)

			fmt.Println()
		}

		// Examples for error handling:
		// - Use helper functions like e.g. errors.IsNotFound()
		// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
		time.Sleep(10 * time.Second)
	}
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
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
