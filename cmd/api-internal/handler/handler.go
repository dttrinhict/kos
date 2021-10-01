package handler

import (
	"github.com/gin-gonic/gin"
	"kos/delivery/rest"
)

type K8sOps struct {
	K8s *rest.K8sOpsRest
}

func NewK8sOps(k8sOpsRest *rest.K8sOpsRest) *K8sOps {
	return &K8sOps{
		K8s: k8sOpsRest,
	}
}

func NewServer(deps *K8sOps) *gin.Engine{
	engine := gin.New()
	internal := engine.Group("/kos/v1/internal")
	internal.GET("/health", deps.K8s.Health)
	internal.POST("/webhook",deps.K8s.Webhook)

	k8s := internal.Group("/k8s")

	{
		pod := k8s.Group("/pod")
		{
			pod.POST("/", deps.K8s.K8sPodsList)
		}
	}
	return engine
}

