package model

import (
	"context"
	"github.com/arioki1/dialogflow-api-and-opentelemetry/srv/v1/request"
)

type (
	Dialogflow        struct{}
	DialogflowUseCase interface {
		DialogflowWebhook(ctx context.Context, req request.DialogflowRequest) (interface{}, int, error)
	}
	DialogflowRepository interface{}
)
