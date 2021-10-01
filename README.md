# kos

Kos là một bộ công cụ dùng đề thực hiện việc tương tác với Kubernetes thông qua api được implement trong nội cụm k8s

#
- Deployment
- Job
- Cronjob
- ...

# Fix Errors:
2021/06/10 12:40:35 Failed to create K8s clientset no Auth Provider found for name "gcp"
```text
https://github.com/kubernetes/client-go/issues/242

It seems you need import the auth plugin package
_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
// or
_ "k8s.io/client-go/plugin/pkg/client/auth"
Is this working as intended? Is there any benefit to not loading auth plugins by default?
```

#In cluster client config
https://github.com/kubernetes/client-go/blob/master/examples/in-cluster-client-configuration/main.go
```text
config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
```

# Forbidden - service account không có quyền

is forbidden: User \"system:serviceaccount:namespace:default\" cannot get resource \"deployments\" in API group \"apps\" in the namespace \"default\""

Để khắc phục lỗi này cần config cho pod một service account có đủ quyền tương tác với k8s
```text
apiVersion: v1
kind: ServiceAccount
metadata:
  name: ${SERVICE_ACCOUNT_NAME}
  namespace: ${NAMESPACE}
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: ${CLUSTER_ROLE_NAME}
rules:
- apiGroups: ["*"]
  resources: ["*"]
  verbs: ["*"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: ${CLUSTER_ROLE_BINDING_NAME}
subjects:
- kind: ServiceAccount
  name: ${SERVICE_ACCOUNT_NAME}
  namespace: ${NAMESPACE}
roleRef:
  kind: ClusterRole
  name: ${CLUSTER_ROLE_NAME}
  apiGroup: rbac.authorization.k8s.io
```


##APIs
#Deployment
POST http://your-domain/kos/v1/internal/k8s/deployment

Request
```json
{
  "metaData": {},
  "deployment": {
    "metadata": {
      "name": "test2",
      "namespace": "default"
    },
    "spec": {
      "replicas": 2,
      "selector": {
        "matchLabels": {
          "app": "test2",
          "name": "test2"
        }
      },
      "template": {
        "metadata": {
          "labels": {
            "name": "test2",
            "app": "test2"
          }
        },
        "spec": {
          "containers": [
            {
              "name": "main",
              "image": "asia.gcr.io/devops/pegasus-api-internal:279073ff",
              "command": [
                "/bin/sh"
              ],
              "args": [
                "-c",
                "/app/main"
              ],
              "resources": {
                "limits": {
                  "cpu": "1500m",
                  "memory": "1Gi"
                },
                "requests": {
                  "cpu": "60m",
                  "memory": "110Mi"
                }
              },
              "livenessProbe": {
                "httpGet": {
                  "path": "/pegasus/v1/internal/health",
                  "port": 8080,
                  "scheme": "HTTP",
                  "httpHeaders": [
                    {
                      "name": "X-Device-ID",
                      "value": "HealthCheck"
                    }
                  ]
                },
                "initialDelaySeconds": 30,
                "timeoutSeconds": 15,
                "periodSeconds": 10,
                "successThreshold": 1,
                "failureThreshold": 3
              },
              "readinessProbe": {
                "httpGet": {
                  "path": "/pegasus/v1/internal/health",
                  "port": 8080,
                  "scheme": "HTTP",
                  "httpHeaders": [
                    {
                      "name": "X-Device-ID",
                      "value": "HealthCheck"
                    }
                  ]
                },
                "initialDelaySeconds": 30,
                "timeoutSeconds": 15,
                "periodSeconds": 10,
                "successThreshold": 1,
                "failureThreshold": 3
              },
              "lifecycle": {
                "preStop": {
                  "exec": {
                    "command": [
                      "/bin/bash",
                      "-c",
                      "sleep 15"
                    ]
                  }
                }
              },
              "terminationMessagePath": "/dev/termination-log",
              "terminationMessagePolicy": "File",
              "imagePullPolicy": "IfNotPresent"
            }
          ],
          "restartPolicy": "Always",
          "terminationGracePeriodSeconds": 30,
          "dnsPolicy": "ClusterFirst",
          "securityContext": {},
          "imagePullSecrets": [
            {
              "name": "docker-image-pull-secret"
            }
          ],
          "affinity": {
            "podAffinity": {
              "requiredDuringSchedulingIgnoredDuringExecution": [
                {
                  "labelSelector": {
                    "matchLabels": {
                      "name": "test4"
                    },
                    "matchExpressions": [
                      {
                        "key": "app",
                        "operator": "In",
                        "values": [
                          "test4",
                          "b",
                          "c",
                          "d"
                        ]
                      }
                    ]
                  },
                  "namespaces": [
                    "default"
                  ],
                  "topologyKey": "topology.kubernetes.io/zone"
                }
              ]
            },
            "podAntiAffinity": {
              "requiredDuringSchedulingIgnoredDuringExecution": [
                {
                  "labelSelector": {
                    "matchLabels": {
                      "name": "test1"
                    },
                    "matchExpressions": [
                      {
                        "key": "app",
                        "operator": "In",
                        "values": [
                          "test1",
                          "b",
                          "c",
                          "d"
                        ]
                      }
                    ]
                  },
                  "namespaces": [
                    "default"
                  ],
                  "topologyKey": "topology.kubernetes.io/zone"
                }
              ]
            }
          },
          "schedulerName": "default-scheduler",
          "tolerations": [
            {
              "key": "preemptible",
              "operator": "Equal",
              "value": "false",
              "effect": "NoSchedule"
            }
          ]
        }
      },
      "strategy": {
        "type": "RollingUpdate",
        "rolling_update": {
          "max_unavailable": 0,
          "max_surge": 1
        }
      },
      "progress_deadline_seconds": 120
    }
  }
}
```

Response
```json
{
    "data": {
        "deployment": {
            "metadata": {
                "name": "test1",
                "namespace": "default",
                "selfLink": "/apis/apps/v1/namespaces/default/deployments/test1",
                "uid": "f5118eef-cbd5-425e-82d5-96f1e8f949ea",
                "resourceVersion": "174000946",
                "generation": 9,
                "creationTimestamp": "2021-06-16T03:44:31Z",
                "managedFields": [
                    {
                        "manager": "main",
                        "operation": "Update",
                        "apiVersion": "apps/v1",
                        "time": "2021-06-16T07:13:05Z",
                        "fieldsType": "FieldsV1",
                        "fieldsV1": {
                            "f:spec": {
                                "f:progressDeadlineSeconds": {},
                                "f:replicas": {},
                                "f:revisionHistoryLimit": {},
                                "f:selector": {
                                    "f:matchLabels": {
                                        ".": {},
                                        "f:app": {},
                                        "f:name": {}
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
                                            "f:app": {},
                                            "f:name": {}
                                        }
                                    },
                                    "f:spec": {
                                        "f:affinity": {
                                            ".": {},
                                            "f:podAffinity": {
                                                ".": {},
                                                "f:requiredDuringSchedulingIgnoredDuringExecution": {}
                                            },
                                            "f:podAntiAffinity": {
                                                ".": {},
                                                "f:requiredDuringSchedulingIgnoredDuringExecution": {}
                                            }
                                        },
                                        "f:containers": {
                                            "k:{\"name\":\"main\"}": {
                                                ".": {},
                                                "f:args": {},
                                                "f:command": {},
                                                "f:image": {},
                                                "f:imagePullPolicy": {},
                                                "f:lifecycle": {
                                                    ".": {},
                                                    "f:preStop": {
                                                        ".": {},
                                                        "f:exec": {
                                                            ".": {},
                                                            "f:command": {}
                                                        }
                                                    }
                                                },
                                                "f:livenessProbe": {
                                                    ".": {},
                                                    "f:failureThreshold": {},
                                                    "f:httpGet": {
                                                        ".": {},
                                                        "f:httpHeaders": {},
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
                                                "f:readinessProbe": {
                                                    ".": {},
                                                    "f:failureThreshold": {},
                                                    "f:httpGet": {
                                                        ".": {},
                                                        "f:httpHeaders": {},
                                                        "f:path": {},
                                                        "f:port": {},
                                                        "f:scheme": {}
                                                    },
                                                    "f:initialDelaySeconds": {},
                                                    "f:periodSeconds": {},
                                                    "f:successThreshold": {},
                                                    "f:timeoutSeconds": {}
                                                },
                                                "f:resources": {
                                                    ".": {},
                                                    "f:limits": {
                                                        ".": {},
                                                        "f:cpu": {},
                                                        "f:memory": {}
                                                    },
                                                    "f:requests": {
                                                        ".": {},
                                                        "f:cpu": {},
                                                        "f:memory": {}
                                                    }
                                                },
                                                "f:terminationMessagePath": {},
                                                "f:terminationMessagePolicy": {}
                                            }
                                        },
                                        "f:dnsPolicy": {},
                                        "f:imagePullSecrets": {
                                            ".": {},
                                            "k:{\"name\":\"docker-image-pull-secret\"}": {
                                                ".": {},
                                                "f:name": {}
                                            }
                                        },
                                        "f:restartPolicy": {},
                                        "f:schedulerName": {},
                                        "f:securityContext": {},
                                        "f:terminationGracePeriodSeconds": {},
                                        "f:tolerations": {}
                                    }
                                }
                            }
                        }
                    }
                ]
            },
            "spec": {
                "replicas": 2,
                "selector": {
                    "matchLabels": {
                        "app": "test1",
                        "name": "test1"
                    }
                },
                "template": {
                    "metadata": {
                        "creationTimestamp": null,
                        "labels": {
                            "app": "test1",
                            "name": "test1"
                        }
                    },
                    "spec": {
                        "containers": [
                            {
                                "name": "main",
                                "image": "asia.gcr.io/devops/pegasus-api-internal:279073ff",
                                "command": [
                                    "/bin/sh"
                                ],
                                "args": [
                                    "-c",
                                    "/app/main"
                                ],
                                "resources": {
                                    "limits": {
                                        "cpu": "1500m",
                                        "memory": "1Gi"
                                    },
                                    "requests": {
                                        "cpu": "60m",
                                        "memory": "110Mi"
                                    }
                                },
                                "livenessProbe": {
                                    "httpGet": {
                                        "path": "/pegasus/v1/internal/health",
                                        "port": 8080,
                                        "scheme": "HTTP",
                                        "httpHeaders": [
                                            {
                                                "name": "X-Device-ID",
                                                "value": "HealthCheck"
                                            }
                                        ]
                                    },
                                    "initialDelaySeconds": 30,
                                    "timeoutSeconds": 15,
                                    "periodSeconds": 10,
                                    "successThreshold": 1,
                                    "failureThreshold": 3
                                },
                                "readinessProbe": {
                                    "httpGet": {
                                        "path": "/pegasus/v1/internal/health",
                                        "port": 8080,
                                        "scheme": "HTTP",
                                        "httpHeaders": [
                                            {
                                                "name": "X-Device-ID",
                                                "value": "HealthCheck"
                                            }
                                        ]
                                    },
                                    "initialDelaySeconds": 30,
                                    "timeoutSeconds": 15,
                                    "periodSeconds": 10,
                                    "successThreshold": 1,
                                    "failureThreshold": 3
                                },
                                "lifecycle": {
                                    "preStop": {
                                        "exec": {
                                            "command": [
                                                "/bin/bash",
                                                "-c",
                                                "sleep 15"
                                            ]
                                        }
                                    }
                                },
                                "terminationMessagePath": "/dev/termination-log",
                                "terminationMessagePolicy": "File",
                                "imagePullPolicy": "IfNotPresent"
                            }
                        ],
                        "restartPolicy": "Always",
                        "terminationGracePeriodSeconds": 30,
                        "dnsPolicy": "ClusterFirst",
                        "securityContext": {},
                        "imagePullSecrets": [
                            {
                                "name": "docker-image-pull-secret"
                            }
                        ],
                        "affinity": {
                            "podAffinity": {
                                "requiredDuringSchedulingIgnoredDuringExecution": [
                                    {
                                        "labelSelector": {
                                            "matchLabels": {
                                                "name": "test4"
                                            },
                                            "matchExpressions": [
                                                {
                                                    "key": "app",
                                                    "operator": "In",
                                                    "values": [
                                                        "test4",
                                                        "b",
                                                        "c",
                                                        "d"
                                                    ]
                                                }
                                            ]
                                        },
                                        "namespaces": [
                                            "default"
                                        ],
                                        "topologyKey": "topology.kubernetes.io/zone"
                                    }
                                ]
                            },
                            "podAntiAffinity": {
                                "requiredDuringSchedulingIgnoredDuringExecution": [
                                    {
                                        "labelSelector": {
                                            "matchLabels": {
                                                "name": "test1"
                                            },
                                            "matchExpressions": [
                                                {
                                                    "key": "app",
                                                    "operator": "In",
                                                    "values": [
                                                        "test1",
                                                        "b",
                                                        "c",
                                                        "d"
                                                    ]
                                                }
                                            ]
                                        },
                                        "namespaces": [
                                            "default"
                                        ],
                                        "topologyKey": "topology.kubernetes.io/zone"
                                    }
                                ]
                            }
                        },
                        "schedulerName": "default-scheduler",
                        "tolerations": [
                            {
                                "key": "preemptible",
                                "operator": "Equal",
                                "value": "false",
                                "effect": "NoSchedule"
                            }
                        ]
                    }
                },
                "strategy": {
                    "type": "RollingUpdate",
                    "rollingUpdate": {
                        "maxUnavailable": "25%",
                        "maxSurge": "25%"
                    }
                },
                "revisionHistoryLimit": 10,
                "progressDeadlineSeconds": 600
            },
            "status": {
                "observedGeneration": 8,
                "replicas": 3,
                "updatedReplicas": 1,
                "readyReplicas": 2,
                "availableReplicas": 2,
                "unavailableReplicas": 1,
                "conditions": [
                    {
                        "type": "Available",
                        "status": "True",
                        "lastUpdateTime": "2021-06-16T07:13:40Z",
                        "lastTransitionTime": "2021-06-16T07:13:40Z",
                        "reason": "MinimumReplicasAvailable",
                        "message": "Deployment has minimum availability."
                    },
                    {
                        "type": "Progressing",
                        "status": "False",
                        "lastUpdateTime": "2021-06-16T07:33:24Z",
                        "lastTransitionTime": "2021-06-16T07:33:24Z",
                        "reason": "ProgressDeadlineExceeded",
                        "message": "ReplicaSet \"test1-5cd874dd7b\" has timed out progressing."
                    }
                ]
            }
        }
    },
    "meta": {
        "code": 202,
        "request_id": "b38b4106-5e47-4238-865e-0e893170396c"
    }
}
```
#Cronjob
POST http://your-domain/kos/v1/internal/k8s/cron-job

Request Body
```json
{
  "metaData": {},
  "cronJob": {
    "metadata": {
      "name": "cron1",
      "namespace": "default"
    },
    "spec": {
      "schedule": "*/3 * * * *",
      "concurrencyPolicy": "Forbid",
      "suspend": false,
      "jobTemplate": {
        "metadata": {
          "name": "cron1",
          "creationTimestamp": null,
          "labels": {
            "app": "cron1",
            "name": "cron1"
          }
        },
        "spec": {
          "template": {
            "metadata": {
              "name": "cron1",
              "creationTimestamp": null,
              "labels": {
                "app": "cron1",
                "name": "cron1"
              },
              "annotations": {
                "sidecar.istio.io/inject": "false"
              }
            },
            "spec": {
              "containers": [
                {
                  "name": "main",
                  "image": "busybox",
                  "command": [
                    "/bin/bash",
                    "-c",
                    "echo 2"
                  ],
                  "resources": {},
                  "terminationMessagePath": "/dev/termination-log",
                  "terminationMessagePolicy": "File",
                  "imagePullPolicy": "Always"
                }
              ],
              "restartPolicy": "Never",
              "terminationGracePeriodSeconds": 30,
              "dnsPolicy": "ClusterFirst",
              "securityContext": {},
              "schedulerName": "default-scheduler"
            }
          }
        }
      },
      "successfulJobsHistoryLimit": 3,
      "failedJobsHistoryLimit": 1
    }
  }
}
```

Response

```json
{
    "data": {
        "cronjob": {
            "metadata": {
                "name": "cron1",
                "namespace": "default",
                "selfLink": "/apis/batch/v1beta1/namespaces/default/cronjobs/cron1",
                "uid": "d9076b6e-eb0b-46cb-842d-d4f516628f1d",
                "resourceVersion": "180619124",
                "creationTimestamp": "2021-06-23T08:50:47Z",
                "managedFields": [
                    {
                        "manager": "main",
                        "operation": "Update",
                        "apiVersion": "batch/v1beta1",
                        "time": "2021-06-23T08:50:47Z",
                        "fieldsType": "FieldsV1",
                        "fieldsV1": {
                            "f:spec": {
                                "f:concurrencyPolicy": {},
                                "f:failedJobsHistoryLimit": {},
                                "f:jobTemplate": {
                                    "f:metadata": {
                                        "f:labels": {
                                            ".": {},
                                            "f:app": {},
                                            "f:name": {}
                                        },
                                        "f:name": {}
                                    },
                                    "f:spec": {
                                        "f:template": {
                                            "f:metadata": {
                                                "f:annotations": {
                                                    ".": {},
                                                    "f:sidecar.istio.io/inject": {}
                                                },
                                                "f:labels": {
                                                    ".": {},
                                                    "f:app": {},
                                                    "f:name": {}
                                                },
                                                "f:name": {}
                                            },
                                            "f:spec": {
                                                "f:containers": {
                                                    "k:{\"name\":\"main\"}": {
                                                        ".": {},
                                                        "f:command": {},
                                                        "f:image": {},
                                                        "f:imagePullPolicy": {},
                                                        "f:name": {},
                                                        "f:resources": {},
                                                        "f:terminationMessagePath": {},
                                                        "f:terminationMessagePolicy": {}
                                                    }
                                                },
                                                "f:dnsPolicy": {},
                                                "f:restartPolicy": {},
                                                "f:schedulerName": {},
                                                "f:securityContext": {},
                                                "f:terminationGracePeriodSeconds": {}
                                            }
                                        }
                                    }
                                },
                                "f:schedule": {},
                                "f:successfulJobsHistoryLimit": {},
                                "f:suspend": {}
                            }
                        }
                    }
                ]
            },
            "spec": {
                "schedule": "*/10 * * * *",
                "concurrencyPolicy": "Forbid",
                "suspend": false,
                "jobTemplate": {
                    "metadata": {
                        "name": "cron1",
                        "creationTimestamp": null,
                        "labels": {
                            "app": "cron1",
                            "name": "cron1"
                        }
                    },
                    "spec": {
                        "template": {
                            "metadata": {
                                "name": "cron1",
                                "creationTimestamp": null,
                                "labels": {
                                    "app": "cron1",
                                    "name": "cron1"
                                },
                                "annotations": {
                                    "sidecar.istio.io/inject": "false"
                                }
                            },
                            "spec": {
                                "containers": [
                                    {
                                        "name": "main",
                                        "image": "busybox",
                                        "command": [
                                            "/bin/bash",
                                            "-c",
                                            "echo 1"
                                        ],
                                        "resources": {},
                                        "terminationMessagePath": "/dev/termination-log",
                                        "terminationMessagePolicy": "File",
                                        "imagePullPolicy": "Always"
                                    }
                                ],
                                "restartPolicy": "Never",
                                "terminationGracePeriodSeconds": 30,
                                "dnsPolicy": "ClusterFirst",
                                "securityContext": {},
                                "schedulerName": "default-scheduler"
                            }
                        }
                    }
                },
                "successfulJobsHistoryLimit": 3,
                "failedJobsHistoryLimit": 1
            },
            "status": {}
        }
    },
    "meta": {
        "code": 202,
        "request_id": "f72f8092-a78a-4ba6-8645-822976f4d2ce"
    }
}
```
#Configmap

POST: {{URL}}/kos/v1/internal/k8s/configmap

Request Body:
```json
{
    "metaData": {},
    "configMap": {
        "metadata": {
            "name": "test23",
            "namespace": "default"
        },
        "data": {
            "a": "a",
            "b": "b"
        }
    }
}
```

Data Response:
```json
{
    "data": {
        "ConfigMap": {
            "metadata": {
                "name": "test23",
                "namespace": "default",
                "selfLink": "/api/v1/namespaces/default/configmaps/test23",
                "uid": "39483999-61ec-4de4-9d99-472eb6062a9d",
                "resourceVersion": "181375576",
                "creationTimestamp": "2021-06-24T04:05:02Z",
                "managedFields": [
                    {
                        "manager": "main",
                        "operation": "Update",
                        "apiVersion": "v1",
                        "time": "2021-06-24T04:05:02Z",
                        "fieldsType": "FieldsV1",
                        "fieldsV1": {
                            "f:data": {
                                ".": {},
                                "f:a": {},
                                "f:b": {}
                            }
                        }
                    }
                ]
            },
            "data": {
                "a": "a",
                "b": "b"
            }
        }
    },
    "meta": {
        "code": 202,
        "request_id": "5e0372b4-93ee-45cb-9916-deb552be8a38"
    }
}
```

#Secret

POST: {{URL}}/kos/v1/internal/k8s/secret

Request Body:

```json
{
    "metaData": {},
    "secret": {
        "metadata": {
            "name": "test23",
            "namespace": "default"
        },
        "data": {
            "a": "ZHNhZGZhc2ZzYWpraGtqaGtqaGtqaGtqaGtq"
        }
    }
}
```

Data Response:
```json
{
    "data": {
        "Secret": {
            "metadata": {
                "name": "test23",
                "namespace": "default",
                "selfLink": "/api/v1/namespaces/default/secrets/test23",
                "uid": "90cc20c5-ff71-4959-b0a0-80967fdc9674",
                "resourceVersion": "181461536",
                "creationTimestamp": "2021-06-24T06:12:41Z",
                "managedFields": [
                    {
                        "manager": "main",
                        "operation": "Update",
                        "apiVersion": "v1",
                        "time": "2021-06-24T06:12:41Z",
                        "fieldsType": "FieldsV1",
                        "fieldsV1": {
                            "f:data": {
                                ".": {},
                                "f:a": {}
                            },
                            "f:type": {}
                        }
                    }
                ]
            },
            "data": {
                "a": "ZHNhZGZhc2ZzYWpraGtqaGtqaGtqaGtqaGtq"
            },
            "type": "Opaque"
        }
    },
    "meta": {
        "code": 202,
        "request_id": "6e73cb39-1eae-49fc-aca9-ec1e529ee6fe"
    }
}
```
#Service
POST: {{URL}}/kos/v1/internal/k8s/service

Request Body:

```json
{
    "metaData": {},
    "service": {
        "metadata": {
            "labels": {
                "name": "test23"
            },
            "name": "test23",
            "namespace": "default"
        },
        "spec": {
            "ports": [
                {
                    "name": "http",
                    "port": 80,
                    "protocol": "TCP",
                    "targetPort": 4445
                }
            ],
            "selector": {
                "name": "label selector of pod"
            },
            "type": "ClusterIP"
        }
    }
}
```

Data Response:
```json
{
    "data": {
        "service": {
            "metadata": {
                "name": "test23",
                "namespace": "default",
                "selfLink": "/api/v1/namespaces/default/services/test23",
                "uid": "5359032a-248e-4f53-8f04-2cc103703a76",
                "resourceVersion": "181556860",
                "creationTimestamp": "2021-06-24T08:10:52Z",
                "labels": {
                    "name": "test23"
                },
                "managedFields": [
                    {
                        "manager": "main",
                        "operation": "Update",
                        "apiVersion": "v1",
                        "time": "2021-06-24T08:10:52Z",
                        "fieldsType": "FieldsV1",
                        "fieldsV1": {
                            "f:metadata": {
                                "f:labels": {
                                    ".": {},
                                    "f:name": {}
                                }
                            },
                            "f:spec": {
                                "f:ports": {
                                    ".": {},
                                    "k:{\"port\":80,\"protocol\":\"TCP\"}": {
                                        ".": {},
                                        "f:name": {},
                                        "f:port": {},
                                        "f:protocol": {},
                                        "f:targetPort": {}
                                    }
                                },
                                "f:selector": {
                                    ".": {},
                                    "f:name": {}
                                },
                                "f:sessionAffinity": {},
                                "f:type": {}
                            }
                        }
                    }
                ]
            },
            "spec": {
                "ports": [
                    {
                        "name": "http",
                        "protocol": "TCP",
                        "port": 80,
                        "targetPort": 4445
                    }
                ],
                "selector": {
                    "name": "label selector of pod"
                },
                "clusterIP": "172.24.60.143",
                "type": "ClusterIP",
                "sessionAffinity": "None"
            },
            "status": {
                "loadBalancer": {}
            }
        }
    },
    "meta": {
        "code": 202,
        "request_id": "14e574da-d07d-47e0-8df6-b2db367fc9df"
    }
}
```
#Istio Virtual Service

POST:

Request Body:
```json
{
    "metaData": {},
    "virtualService": {
        "metadata": {
            "name": "test1",
            "namespace": "default"
        },
        "spec": {
            "gateways": [
                "internal-gateway-tls.istio-system.svc.cluster.local"
            ],
            "hosts": [
                "api-uat.domain.dev",
                "service-name"
            ],
            "http": [
                {
                    "match": [
                        {
                            "uri": {
                                "prefix": "/test/v1"
                            }
                        },
                        {
                            "uri": {
                                "prefix": "/test/v2"
                            }
                        },
                        {
                            "uri": {
                                "prefix": "/test/v1/me/profile"
                            }
                        },
                        {
                            "uri": {
                                "prefix": "/test/swagger"
                            }
                        }
                    ],
                    "route": [
                        {
                            "destination": {
                                "host": "service-name-v1"
                            },
                            "weight": 70
                        },
                        {
                            "destination": {
                                "host": "service-name-v2"
                            },
                            "weight": 30
                        }
                    ]
                }
            ]
        }
    }
}
```

Data Response:

```json
{
    "data": {
        "VirtualService": {
            "metadata": {
                "name": "test1",
                "namespace": "default",
                "selfLink": "/apis/networking.istio.io/v1beta1/namespaces/default/virtualservices/test1",
                "uid": "7893b3a6-3ee2-4c88-9a6e-f6a2754cfe55",
                "resourceVersion": "181630230",
                "generation": 1,
                "creationTimestamp": "2021-06-24T10:23:45Z",
                "managedFields": [
                    {
                        "manager": "main",
                        "operation": "Update",
                        "apiVersion": "networking.istio.io/v1beta1",
                        "time": "2021-06-24T10:23:45Z",
                        "fieldsType": "FieldsV1",
                        "fieldsV1": {
                            "f:spec": {
                                ".": {},
                                "f:gateways": {},
                                "f:hosts": {},
                                "f:http": {}
                            },
                            "f:status": {}
                        }
                    }
                ]
            },
            "spec": {
                "hosts": [
                    "api-uat.domain.dev",
                    "service-name"
                ],
                "gateways": [
                    "internal-gateway-tls.istio-system.svc.cluster.local"
                ],
                "http": [
                    {
                        "match": [
                            {
                                "uri": {
                                    "prefix": "/test/v1"
                                }
                            },
                            {
                                "uri": {
                                    "prefix": "/test/v2"
                                }
                            },
                            {
                                "uri": {
                                    "prefix": "/test/v1/me/profile"
                                }
                            },
                            {
                                "uri": {
                                    "prefix": "/test/swagger"
                                }
                            }
                        ],
                        "route": [
                            {
                                "destination": {
                                    "host": "service-name-v1"
                                },
                                "weight": 70
                            },
                            {
                                "destination": {
                                    "host": "service-name-v2"
                                },
                                "weight": 30
                            }
                        ]
                    }
                ]
            },
            "status": {}
        }
    },
    "meta": {
        "code": 202,
        "request_id": "ab8d7da5-6c83-41ac-a931-f667d4656ac7"
    }
}
```

#HorizontalPodAutoScaler

POST: {{URL}}/kos/v1/internal/k8s/hpa

Request Body
```json
{
    "metaData": {},
    "horizontalPodAutoscaler": {
        "metadata": {
            "name": "test4",
            "namespace": "default"
        },
        "spec": {
            "maxReplicas": 2,
            "minReplicas": 1,
            "scaleTargetRef": {
                "apiVersion": "apps/v1",
                "kind": "Deployment",
                "name": "test4"
            },
            "targetCPUUtilizationPercentage": 80
        }
    }
}
```

Data Response
```json
{
    "data": {
        "horizontalPodAutoscaler": {
            "metadata": {
                "name": "test4",
                "namespace": "default",
                "selfLink": "/apis/autoscaling/v1/namespaces/default/horizontalpodautoscalers/test4",
                "uid": "37d4015f-c77f-4eea-8819-6672997f28cf",
                "resourceVersion": "182291191",
                "creationTimestamp": "2021-06-25T03:11:52Z",
                "managedFields": [
                    {
                        "manager": "main",
                        "operation": "Update",
                        "apiVersion": "autoscaling/v1",
                        "time": "2021-06-25T03:11:52Z",
                        "fieldsType": "FieldsV1",
                        "fieldsV1": {
                            "f:spec": {
                                "f:maxReplicas": {},
                                "f:minReplicas": {},
                                "f:scaleTargetRef": {
                                    "f:apiVersion": {},
                                    "f:kind": {},
                                    "f:name": {}
                                },
                                "f:targetCPUUtilizationPercentage": {}
                            }
                        }
                    }
                ]
            },
            "spec": {
                "scaleTargetRef": {
                    "kind": "Deployment",
                    "name": "test4",
                    "apiVersion": "apps/v1"
                },
                "minReplicas": 1,
                "maxReplicas": 2,
                "targetCPUUtilizationPercentage": 80
            },
            "status": {
                "currentReplicas": 0,
                "desiredReplicas": 0
            }
        }
    },
    "meta": {
        "code": 202,
        "request_id": "bd3ac37c-27a8-4a55-af99-c09eac6c0fca"
    }
}
```

#Job
POST:

Request Body:
```json
{
    "metaData": {},
    "job": {
        "metadata": {
        "labels": {
            "job-name": "job"
        },
        "name": "job",
        "namespace": "default"
        },
        "spec": {
            "backoffLimit": 0,
            "completions": 1,
            "parallelism": 1,
            "selector": {
                "matchLabels": {
                    "job-name": "job"
                }
            },
            "template": {
                "metadata": {
                    "annotations": {
                        "sidecar.istio.io/inject": "false"
                    },
                    "labels": {
                        "job-name": "job"
                    }
                },
                "spec": {
                    "containers": [
                        {
                            "command": [
                                "ls"
                            ],
                            "image": "alpine:latest",
                            "imagePullPolicy": "Always",
                            "name": "job",
                            "resources": {}
                        }
                    ],
                    "dnsPolicy": "ClusterFirst",
                    "restartPolicy": "Never",
                    "schedulerName": "default-scheduler",
                    "securityContext": {},
                    "terminationGracePeriodSeconds": 30
                }
            }
        }
    }
}
```
Data Response:
```json
{
    "data": {
        "job": {
            "metadata": {
                "name": "job",
                "namespace": "default",
                "selfLink": "/apis/batch/v1/namespaces/default/jobs/job",
                "uid": "7ce575c5-c3bf-4d28-88a4-d468fb6bfefc",
                "resourceVersion": "182359587",
                "creationTimestamp": "2021-06-25T04:44:37Z",
                "labels": {
                    "job-name": "job"
                },
                "managedFields": [
                    {
                        "manager": "main",
                        "operation": "Update",
                        "apiVersion": "batch/v1",
                        "time": "2021-06-25T04:44:37Z",
                        "fieldsType": "FieldsV1",
                        "fieldsV1": {
                            "f:metadata": {
                                "f:labels": {
                                    ".": {},
                                    "f:job-name": {}
                                }
                            },
                            "f:spec": {
                                "f:backoffLimit": {},
                                "f:completions": {},
                                "f:parallelism": {},
                                "f:template": {
                                    "f:metadata": {
                                        "f:labels": {
                                            ".": {},
                                            "f:job-name": {}
                                        }
                                    },
                                    "f:spec": {
                                        "f:containers": {
                                            "k:{\"name\":\"job\"}": {
                                                ".": {},
                                                "f:command": {},
                                                "f:image": {},
                                                "f:imagePullPolicy": {},
                                                "f:name": {},
                                                "f:resources": {},
                                                "f:terminationMessagePath": {},
                                                "f:terminationMessagePolicy": {}
                                            }
                                        },
                                        "f:dnsPolicy": {},
                                        "f:restartPolicy": {},
                                        "f:schedulerName": {},
                                        "f:securityContext": {},
                                        "f:terminationGracePeriodSeconds": {}
                                    }
                                }
                            }
                        }
                    }
                ]
            },
            "spec": {
                "parallelism": 1,
                "completions": 1,
                "backoffLimit": 0,
                "selector": {
                    "matchLabels": {
                        "controller-uid": "7ce575c5-c3bf-4d28-88a4-d468fb6bfefc"
                    }
                },
                "template": {
                    "metadata": {
                        "creationTimestamp": null,
                        "labels": {
                            "controller-uid": "7ce575c5-c3bf-4d28-88a4-d468fb6bfefc",
                            "job-name": "job"
                        }
                    },
                    "spec": {
                        "containers": [
                            {
                                "name": "job",
                                "image": "alpine:latest",
                                "command": [
                                    "ls"
                                ],
                                "resources": {},
                                "terminationMessagePath": "/dev/termination-log",
                                "terminationMessagePolicy": "File",
                                "imagePullPolicy": "Always"
                            }
                        ],
                        "restartPolicy": "Never",
                        "terminationGracePeriodSeconds": 30,
                        "dnsPolicy": "ClusterFirst",
                        "securityContext": {},
                        "schedulerName": "default-scheduler"
                    }
                }
            },
            "status": {
                "conditions": [
                    {
                        "type": "Complete",
                        "status": "True",
                        "lastProbeTime": "2021-06-25T04:44:42Z",
                        "lastTransitionTime": "2021-06-25T04:44:42Z"
                    }
                ],
                "startTime": "2021-06-25T04:44:37Z",
                "completionTime": "2021-06-25T04:44:42Z",
                "succeeded": 1
            }
        }
    },
    "meta": {
        "code": 202,
        "request_id": "10dce27f-6275-4cc8-8cfa-d13e891f3701"
    }
}
```

#Pods
Get pods by namespace

POST: {{URL}}/kos/v1/internal/k8s/pod/

Request Body
```json
{
  "namespace": "namespace's name"
}
```
Data response

success
```json
{
  "data": [
    "specs of pods"
  ],
  "meta": {
    "code": 202,
    "request_id": "9fe7bb39-aa60-4f8a-bb8f-71a4433185ba"
  }
}
```
Error
```json
{
    "data": "The request body data is wrong!",
    "meta": {
        "code": 500000,
        "message": "Internal server error",
        "request_id": "2badae66-39bd-46ea-b255-0a9e1f29e770"
    }
}
```

#Cluster Role
POST: {{URL}}/kos/v1/internal/k8s/cluster-role

Request Body
```json
{
    "metaData": {},
    "clusterRole": {
        "metadata": {
        "name": "cr-test"
        },
        "rules": [
            {
                "apiGroups": [
                    ""
                ],
                "resources": [
                    "pods"
                ],
                "verbs": [
                    "get",
                    "wat",
                    "list"
                ]
            }
        ]
    }
}
```
Data Response

Success
```json
{
    "data": {
        "clusterRole": {
            "metadata": {
                "name": "cr-test",
                "selfLink": "/apis/rbac.authorization.k8s.io/v1/clusterroles/cr-test",
                "uid": "0c41c4b5-35d2-41db-bde3-4b4c6c7687e8",
                "resourceVersion": "182591112",
                "creationTimestamp": "2021-06-25T10:34:20Z",
                "managedFields": [
                    {
                        "manager": "main",
                        "operation": "Update",
                        "apiVersion": "rbac.authorization.k8s.io/v1",
                        "time": "2021-06-25T10:34:20Z",
                        "fieldsType": "FieldsV1",
                        "fieldsV1": {
                            "f:rules": {}
                        }
                    }
                ]
            },
            "rules": [
                {
                    "verbs": [
                        "get",
                        "wat",
                        "list"
                    ],
                    "apiGroups": [
                        ""
                    ],
                    "resources": [
                        "pods"
                    ]
                }
            ]
        }
    },
    "meta": {
        "code": 202,
        "request_id": "773d4d58-e46e-40d8-85e2-397cacc43257"
    }
}
```

Error
```json
{
    "data": "Error while create Cluster role component, clusterroles.rbac.authorization.k8s.io \"cr-test\" already exists",
    "meta": {
        "code": 500000,
        "message": "Internal server error",
        "request_id": "f287303a-2c67-4d48-b6db-fb51ba03cc87"
    }
}
```