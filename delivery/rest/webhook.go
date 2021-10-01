package rest

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

func (e *K8sOpsRest) Webhook(ginCtx *gin.Context) {
	jsonData, err := ioutil.ReadAll(ginCtx.Request.Body)
	if err != nil {
		ginCtx.JSON(500, err.Error() + "looix ow day a")
	}
	var reqBody interface{}
	err = json.Unmarshal(jsonData, &reqBody)
	ginCtx.JSON(200, reqBody)
	log.Printf("Webhook Response: %v", reqBody)
}