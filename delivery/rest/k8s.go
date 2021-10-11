package rest

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	usecase "kos/usecases/k8s"
)

const HTTP_STATUS_OK = 200
const HTTP_STATUS_BAD_REQUEST = 400

type K8sOpsRest struct {
	K8sUseCase usecase.Kubernetes
}

func NewK8sOpsRest(k8sUseCase usecase.Kubernetes) *K8sOpsRest {
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
		Response, err := e.K8sUseCase.ListPods(req["namespace"])
		if err != nil {
			return
		}

		ginCtx.JSON(HTTP_STATUS_OK, gin.H{
				"respone": Response,
			})
	}else{
		ginCtx.JSON(HTTP_STATUS_BAD_REQUEST, gin.H{
			"error": "error",
		})
	}
}

func (e *K8sOpsRest) Health(ginCtx *gin.Context)  {
	ginCtx.JSON(HTTP_STATUS_OK, gin.H{
		"response": "running",
	})
}