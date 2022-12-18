package usecase

import (
	"context"
	"github.com/arioki1/dialogflow-api-and-opentelemetry/config"
	"github.com/arioki1/dialogflow-api-and-opentelemetry/srv/v1/model"
	"github.com/arioki1/dialogflow-api-and-opentelemetry/srv/v1/request"
	"github.com/arioki1/dialogflow-api-and-opentelemetry/srv/v1/response"
)

type dialogflow struct {
	config               config.Config
	intentName           string
	queryText            string
	action               string
	request              *request.DialogflowRequest
	response             *response.DialogFlowResponseGenerator
	isHandled            bool
}

func (d *dialogflow) DialogflowWebhook(ctx context.Context, req request.DialogflowRequest) (interface{}, int, error) {
	d.intentName = req.QueryResult.Intent.DisplayName
	d.queryText = req.QueryResult.QueryText
	d.action = req.QueryResult.Action
	d.request = &req
	d.response = &response.DialogFlowResponseGenerator{}
	d.response.SetFulfillmentMessages( "Hello, I'm Dialogflow!")
	return d.response, 0, nil
}

func NewDialogflowUseCase(c config.Config) model.DialogflowUseCase {
	return &dialogflow{
		config:               c,
	}
}
