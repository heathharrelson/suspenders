export const json = `
{
  "metadata": {
    "name": "traefik",
    "namespace": "default",
    "selfLink": "/apis/apps/v1/namespaces/default/deployments/traefik",
    "uid": "fac10fe5-3bac-4c0d-bc7f-1bf96da632dd",
    "resourceVersion": "4919787",
    "generation": 1,
    "creationTimestamp": "2020-08-15T21:38:59Z",
    "labels": {
      "app.kubernetes.io/instance": "traefik",
      "app.kubernetes.io/managed-by": "Helm",
      "app.kubernetes.io/name": "traefik",
      "helm.sh/chart": "traefik-8.12.0"
    },
    "annotations": {
      "deployment.kubernetes.io/revision": "1",
      "meta.helm.sh/release-name": "traefik",
      "meta.helm.sh/release-namespace": "default"
    },
    "managedFields": [
      {
        "manager": "Go-http-client",
        "operation": "Update",
        "apiVersion": "apps/v1",
        "time": "2020-08-15T21:38:59Z",
        "fieldsType": "FieldsV1",
        "fieldsV1": {
          "f:metadata": {
            "f:annotations": {
              ".": {},
              "f:meta.helm.sh/release-name": {},
              "f:meta.helm.sh/release-namespace": {}
            },
            "f:labels": {
              ".": {},
              "f:app.kubernetes.io/instance": {},
              "f:app.kubernetes.io/managed-by": {},
              "f:app.kubernetes.io/name": {},
              "f:helm.sh/chart": {}
            }
          },
          "f:spec": {
            "f:progressDeadlineSeconds": {},
            "f:replicas": {},
            "f:revisionHistoryLimit": {},
            "f:selector": {
              "f:matchLabels": {
                ".": {},
                "f:app.kubernetes.io/instance": {},
                "f:app.kubernetes.io/name": {}
              }
            },
            "f:strategy": {
              "f:rollingUpdate": {
                ".": {},
                "f:maxSurge": {},
                "f:maxUnavailable": {}
              },
              "f:type": {}
            },
            "f:template": {
              "f:metadata": {
                "f:labels": {
                  ".": {},
                  "f:app.kubernetes.io/instance": {},
                  "f:app.kubernetes.io/managed-by": {},
                  "f:app.kubernetes.io/name": {},
                  "f:helm.sh/chart": {}
                }
              },
              "f:spec": {
                "f:containers": {
                  "k:{\\"name\\":\\"traefik\\"}": {
                    ".": {},
                    "f:args": {},
                    "f:image": {},
                    "f:imagePullPolicy": {},
                    "f:livenessProbe": {
                      ".": {},
                      "f:failureThreshold": {},
                      "f:httpGet": {
                        ".": {},
                        "f:path": {},
                        "f:port": {},
                        "f:scheme": {}
                      },
                      "f:initialDelaySeconds": {},
                      "f:periodSeconds": {},
                      "f:successThreshold": {},
                      "f:timeoutSeconds": {}
                    },
                    "f:name": {},
                    "f:ports": {
                      ".": {},
                      "k:{\\"containerPort\\":8000,\\"protocol\\":\\"TCP\\"}": {
                        ".": {},
                        "f:containerPort": {},
                        "f:name": {},
                        "f:protocol": {}
                      },
                      "k:{\\"containerPort\\":8443,\\"protocol\\":\\"TCP\\"}": {
                        ".": {},
                        "f:containerPort": {},
                        "f:name": {},
                        "f:protocol": {}
                      },
                      "k:{\\"containerPort\\":9000,\\"protocol\\":\\"TCP\\"}": {
                        ".": {},
                        "f:containerPort": {},
                        "f:name": {},
                        "f:protocol": {}
                      }
                    },
                    "f:readinessProbe": {
                      ".": {},
                      "f:failureThreshold": {},
                      "f:httpGet": {
                        ".": {},
                        "f:path": {},
                        "f:port": {},
                        "f:scheme": {}
                      },
                      "f:initialDelaySeconds": {},
                      "f:periodSeconds": {},
                      "f:successThreshold": {},
                      "f:timeoutSeconds": {}
                    },
                    "f:resources": {},
                    "f:securityContext": {
                      ".": {},
                      "f:capabilities": {
                        ".": {},
                        "f:drop": {}
                      },
                      "f:readOnlyRootFilesystem": {},
                      "f:runAsGroup": {},
                      "f:runAsNonRoot": {},
                      "f:runAsUser": {}
                    },
                    "f:terminationMessagePath": {},
                    "f:terminationMessagePolicy": {},
                    "f:volumeMounts": {
                      ".": {},
                      "k:{\\"mountPath\\":\\"/data\\"}": {
                        ".": {},
                        "f:mountPath": {},
                        "f:name": {}
                      },
                      "k:{\\"mountPath\\":\\"/tmp\\"}": {
                        ".": {},
                        "f:mountPath": {},
                        "f:name": {}
                      }
                    }
                  }
                },
                "f:dnsPolicy": {},
                "f:restartPolicy": {},
                "f:schedulerName": {},
                "f:securityContext": {
                  ".": {},
                  "f:fsGroup": {}
                },
                "f:serviceAccount": {},
                "f:serviceAccountName": {},
                "f:terminationGracePeriodSeconds": {},
                "f:volumes": {
                  ".": {},
                  "k:{\\"name\\":\\"data\\"}": {
                    ".": {},
                    "f:emptyDir": {},
                    "f:name": {}
                  },
                  "k:{\\"name\\":\\"tmp\\"}": {
                    ".": {},
                    "f:emptyDir": {},
                    "f:name": {}
                  }
                }
              }
            }
          }
        }
      },
      {
        "manager": "kube-controller-manager",
        "operation": "Update",
        "apiVersion": "apps/v1",
        "time": "2020-09-08T00:03:05Z",
        "fieldsType": "FieldsV1",
        "fieldsV1": {
          "f:metadata": {
            "f:annotations": {
              "f:deployment.kubernetes.io/revision": {}
            }
          },
          "f:status": {
            "f:availableReplicas": {},
            "f:conditions": {
              ".": {},
              "k:{\\"type\\":\\"Available\\"}": {
                ".": {},
                "f:lastTransitionTime": {},
                "f:lastUpdateTime": {},
                "f:message": {},
                "f:reason": {},
                "f:status": {},
                "f:type": {}
              },
              "k:{\\"type\\":\\"Progressing\\"}": {
                ".": {},
                "f:lastTransitionTime": {},
                "f:lastUpdateTime": {},
                "f:message": {},
                "f:reason": {},
                "f:status": {},
                "f:type": {}
              }
            },
            "f:observedGeneration": {},
            "f:readyReplicas": {},
            "f:replicas": {},
            "f:updatedReplicas": {}
          }
        }
      }
    ]
  },
  "spec": {
    "replicas": 1,
    "selector": {
      "matchLabels": {
        "app.kubernetes.io/instance": "traefik",
        "app.kubernetes.io/name": "traefik"
      }
    },
    "template": {
      "metadata": {
        "creationTimestamp": null,
        "labels": {
          "app.kubernetes.io/instance": "traefik",
          "app.kubernetes.io/managed-by": "Helm",
          "app.kubernetes.io/name": "traefik",
          "helm.sh/chart": "traefik-8.12.0"
        }
      },
      "spec": {
        "volumes": [
          {
            "name": "data",
            "emptyDir": {}
          },
          {
            "name": "tmp",
            "emptyDir": {}
          }
        ],
        "containers": [
          {
            "name": "traefik",
            "image": "traefik:2.2.8",
            "args": [
              "--global.checknewversion",
              "--global.sendanonymoususage",
              "--entryPoints.traefik.address=:9000/tcp",
              "--entryPoints.web.address=:8000/tcp",
              "--entryPoints.websecure.address=:8443/tcp",
              "--api.dashboard=true",
              "--ping=true",
              "--providers.kubernetescrd",
              "--providers.kubernetesingress"
            ],
            "ports": [
              {
                "name": "traefik",
                "containerPort": 9000,
                "protocol": "TCP"
              },
              {
                "name": "web",
                "containerPort": 8000,
                "protocol": "TCP"
              },
              {
                "name": "websecure",
                "containerPort": 8443,
                "protocol": "TCP"
              }
            ],
            "resources": {},
            "volumeMounts": [
              {
                "name": "data",
                "mountPath": "/data"
              },
              {
                "name": "tmp",
                "mountPath": "/tmp"
              }
            ],
            "livenessProbe": {
              "httpGet": {
                "path": "/ping",
                "port": 9000,
                "scheme": "HTTP"
              },
              "initialDelaySeconds": 10,
              "timeoutSeconds": 2,
              "periodSeconds": 10,
              "successThreshold": 1,
              "failureThreshold": 3
            },
            "readinessProbe": {
              "httpGet": {
                "path": "/ping",
                "port": 9000,
                "scheme": "HTTP"
              },
              "initialDelaySeconds": 10,
              "timeoutSeconds": 2,
              "periodSeconds": 10,
              "successThreshold": 1,
              "failureThreshold": 1
            },
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "IfNotPresent",
            "securityContext": {
              "capabilities": {
                "drop": [
                  "ALL"
                ]
              },
              "runAsUser": 65532,
              "runAsGroup": 65532,
              "runAsNonRoot": true,
              "readOnlyRootFilesystem": true
            }
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 60,
        "dnsPolicy": "ClusterFirst",
        "serviceAccountName": "traefik",
        "serviceAccount": "traefik",
        "securityContext": {
          "fsGroup": 65532
        },
        "schedulerName": "default-scheduler"
      }
    },
    "strategy": {
      "type": "RollingUpdate",
      "rollingUpdate": {
        "maxUnavailable": 1,
        "maxSurge": 1
      }
    },
    "revisionHistoryLimit": 10,
    "progressDeadlineSeconds": 600
  },
  "status": {
    "observedGeneration": 1,
    "replicas": 1,
    "updatedReplicas": 1,
    "readyReplicas": 1,
    "availableReplicas": 1,
    "conditions": [
      {
        "type": "Available",
        "status": "True",
        "lastUpdateTime": "2020-08-15T21:39:00Z",
        "lastTransitionTime": "2020-08-15T21:39:00Z",
        "reason": "MinimumReplicasAvailable",
        "message": "Deployment has minimum availability."
      },
      {
        "type": "Progressing",
        "status": "True",
        "lastUpdateTime": "2020-08-15T21:39:36Z",
        "lastTransitionTime": "2020-08-15T21:38:59Z",
        "reason": "NewReplicaSetAvailable",
        "message": "ReplicaSet \\"traefik-68cd8579f5\\" has successfully progressed."
      }
    ]
  }
}
`

export const jsonArray = `[${json}]`
export const deployment = JSON.parse(json)
export const deploymentArray = JSON.parse(jsonArray)
