package registry

import (
	"github.com/arioki1/dialogflow-api-and-opentelemetry/config"
	"sync"
)

type repositoryRegistry struct {
	cfg config.Config
}

type RepositoryRegistry interface {
}

func NewRepositoryRegistry(cfg config.Config) RepositoryRegistry {
	var repoRegistry RepositoryRegistry
	var loadOne sync.Once

	loadOne.Do(func() {
		repoRegistry = &repositoryRegistry{
			cfg: cfg,
		}
	})

	return repoRegistry
}
