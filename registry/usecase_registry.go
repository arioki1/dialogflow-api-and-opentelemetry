package registry

import (
	"github.com/arioki1/dialogflow-api-and-opentelemetry/config"
	"github.com/arioki1/dialogflow-api-and-opentelemetry/srv/v1/model"
	"github.com/arioki1/dialogflow-api-and-opentelemetry/srv/v1/usecase"
	"sync"
)

type useCaseRegistry struct {
	repo RepositoryRegistry
	cfg  config.Config
}

type UseCaseRegistry interface {
	Dialogflow() model.DialogflowUseCase
}

func NewUseCaseRegistry(repo RepositoryRegistry, cfg config.Config) UseCaseRegistry {
	var uc UseCaseRegistry
	var loadOne sync.Once
	loadOne.Do(func() {
		uc = &useCaseRegistry{
			repo: repo,
			cfg:  cfg,
		}
	})

	return uc
}

func (u useCaseRegistry) Dialogflow() model.DialogflowUseCase {
	var qu model.DialogflowUseCase
	var loadOne sync.Once

	loadOne.Do(func() {
		qu = usecase.NewDialogflowUseCase(u.cfg)
	})

	return qu
}
