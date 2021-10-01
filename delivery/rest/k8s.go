package rest

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	usecase "kos/usecases/k8s"
)
type K8sOpsRest struct {
	K8sUseCase usecase.K8sUseCase
}

func NewK8sOpsRest(k8sUseCase usecase.K8sUseCase) *K8sOpsRest {
	return &K8sOpsRest{
		K8sUseCase: k8sUseCase,
	}
}

func (e *K8sOpsRest) K8sPodsList(ginCtx *gin.Context) {
	var req map[string]string
	err := json.NewDecoder(ginCtx.Request.Body).Decode(&req)
	if err != nil {

		return
	}
	if _, ok := req["namespace"]; ok {
		Response, err := e.K8sUseCase.K8sPodsList(req["namespace"])
		if err != nil {
			return
		}
			ginCtx.JSON(200, gin.H{
				"respone": Response,
			})
	}else{
		ginCtx.JSON(503, gin.H{
			"error": "error",
		})
	}
}

func (e *K8sOpsRest) Health(ginCtx *gin.Context)  {
	ginCtx.JSON(200, gin.H{
		"response": "running",
	})
}