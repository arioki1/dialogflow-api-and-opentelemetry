package server

import (
	"context"
	"fmt"
	"github.com/arioki1/dialogflow-api-and-opentelemetry/config"
	"github.com/arioki1/dialogflow-api-and-opentelemetry/registry"
	"github.com/arioki1/dialogflow-api-and-opentelemetry/srv/v1/delivery"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Start(cfg config.Config, router *gin.Engine) {
	log.Info().Msg("starting server..")

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.GetPort()),
		Handler: router,
	}
	go func() {
		initServer(cfg, router)
	}()

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Msgf("listen: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info().Msg("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Msgf("server forced to shutdown %s", err.Error())
	}

	log.Info().Msg("server exiting")
}

func initServer(cfg config.Config, router *gin.Engine) {
	rep := registry.NewRepositoryRegistry(cfg)
	uc := registry.NewUseCaseRegistry(rep, cfg)

	//Dialogflow Delivery
	queueDelivery := delivery.NewDialogflowDelivery(uc.Dialogflow())
	queueGroup := router.Group("/api/v1/webhook")
	queueDelivery.Mount(queueGroup)
}
