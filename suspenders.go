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

Derived from https://github.com/kubernetes/sample-controller/blob/master/main.go
*/

package main

import (
	"flag"
	"net/url"
	"time"

	"github.com/heathharrelson/suspenders/buildinfo"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"

	// Uncomment the following line to load the gcp plugin (only required to authenticate against GKE clusters).
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	"k8s.io/sample-controller/pkg/signals"
)

var (
	externalURL string
	masterURL   string
	kubeconfig  string
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	klog.Infof("Starting suspenders %s", buildinfo.Version)
	klog.Infof(buildinfo.Print())

	// set up signals so we handle the first shutdown signal gracefully
	stopCh := signals.SetupSignalHandler()

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		klog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(kubeClient, time.Second*30)

	extURL, err := url.Parse(externalURL)
	if err != nil {
		klog.Fatalf("Could not parse external URL: %s", err.Error())
	}

	controller := NewServer(kubeClient, kubeInformerFactory.Apps().V1().Deployments(), extURL)

	// notice that there is no need to run Start methods in a separate goroutine. (i.e. go kubeInformerFactory.Start(stopCh)
	// Start method is non-blocking and runs all registered informers in a dedicated goroutine.
	kubeInformerFactory.Start(stopCh)

	if err = controller.Run(stopCh); err != nil {
		klog.Fatalf("Error running controller: %s", err.Error())
	}
}

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&externalURL, "external-url", "", "The URL at which the dashboard will be reachable (e.g. when behind a reverse proxy or ingress). Used to generate links to assets.")
}
