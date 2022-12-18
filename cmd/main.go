package main

import (
	"github.com/arioki1/dialogflow-api-and-opentelemetry/config"
	"github.com/arioki1/dialogflow-api-and-opentelemetry/router"
	"github.com/arioki1/dialogflow-api-and-opentelemetry/server"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {
	_ = godotenv.Load()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Msgf("failed to run server: %s", err.Error())
	}

	if !cfg.GetDebug() {
		gin.SetMode(gin.ReleaseMode)
	}

	routes := router.NewRouter(cfg)
	server.Start(cfg, routes)

}
