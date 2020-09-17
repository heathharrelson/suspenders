# Multi-Arch Build

## Ambassador

Repo: https://github.com/datawire/ambassador

Not applicable; only compiled for amd64.

https://github.com/datawire/ambassador/issues/2121

## Argo CD

Repo: https://github.com/argoproj/argo-cd/

Not applicable; only compiled for amd64.

## Kubernetes Dashboard

Repo: https://github.com/kubernetes/dashboard
Gulp tasks: https://github.com/kubernetes/dashboard/blob/master/aio/gulp/backend.js
Dockerfile: https://github.com/kubernetes/dashboard/blob/master/aio/Dockerfile

* Con: Build is all done with Gulp

## Kube RBAC Proxy

Repo: https://github.com/brancz/kube-rbac-proxy
Makefile: https://github.com/brancz/kube-rbac-proxy/blob/master/Makefile
Dockerfile: https://github.com/brancz/kube-rbac-proxy/blob/master/Dockerfile

### Notes

* Pro: One of the simplest examples; Makefile is not too abstract
* Con: Uses [manifest-tool](https://github.com/estesp/manifest-tool) to build the multi-arch manifest instead of Docker client

## Kubewatch

Repo: https://github.com/bitnami-labs/kubewatch

Not applicable; only compiled for amd64.

## Prometheus

Repo: https://github.com/prometheus/prometheus/
Makefile: https://github.com/prometheus/prometheus/blob/master/Makefile.common
Dockerfile: https://github.com/prometheus/prometheus/blob/master/Dockerfile

### Notes

* Con: Complex Makefile with lots of abstraction and magic targets
* Pro: Straightforward example of `docker manifest` use: https://github.com/prometheus/prometheus/blob/master/Makefile.common#L255

## Prometheus Operator

Repo: https://github.com/prometheus-operator/prometheus-operator
Makefile: https://github.com/prometheus-operator/prometheus-operator/blob/master/Makefile
Dockerfile: https://github.com/prometheus-operator/prometheus-operator/blob/master/Dockerfile
Travis build: https://github.com/prometheus-operator/prometheus-operator/blob/master/scripts/travis-push-docker-image.sh

### Notes

* Pro: Travis build script has clear example of how `docker manifest` command works
* Con: Makefile is kind of complicated

## Traefik

Repo: https://github.com/containous/traefik
Script: https://github.com/containous/traefik/blob/master/script/crossbinary-default

### Notes

* Script linked above builds all the different architectures, but the actual image is actually constructed by Docker Hub: https://github.com/containous/traefik/issues/4291

## Weave Net

Repo: https://github.com/weaveworks/weave
Makefile: https://github.com/weaveworks/weave/blob/master/Makefile
Dockerfile: https://github.com/weaveworks/weave/blob/master/build/Dockerfile

### Notes

* Con: Not a good example. Complex cross-compilation due to use of `gcc`.
* Con: Uses `manifest-tool`.
