package delivery

import (
	"github.com/arioki1/dialogflow-api-and-opentelemetry/helpers"
	"github.com/arioki1/dialogflow-api-and-opentelemetry/srv/v1/model"
	"github.com/arioki1/dialogflow-api-and-opentelemetry/srv/v1/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

type dialogflowDelivery struct {
	dialogflowUC model.DialogflowUseCase
}
type DialogflowDelivery interface {
	Mount(group *gin.RouterGroup)
}

func (q *dialogflowDelivery) Mount(group *gin.RouterGroup) {
	group.POST("/dialogflow", q.DialogflowWebhook)
}

func NewDialogflowDelivery(dialogflowUC model.DialogflowUseCase) DialogflowDelivery {
	return &dialogflowDelivery{
		dialogflowUC: dialogflowUC,
	}
}

func (q *dialogflowDelivery) DialogflowWebhook(c *gin.Context) {
	var requestWebhook request.DialogflowRequest
	if err := c.Bind(&requestWebhook); err != nil {
		helpers.PrintErrStringLog(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"message": "Invalid JSON body",
		})
		return
	}
	res, _, err := q.dialogflowUC.DialogflowWebhook(c, requestWebhook)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "error",
			"error":   true,
		})
		return
	}
	c.JSON(http.StatusOK, res)
}
