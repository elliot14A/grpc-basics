package client

import (
	"log"

	"github.com/elliot14A/tcorp/assessment/client/app"
	"github.com/elliot14A/tcorp/assessment/client/internal/infrastructure/http"
	"github.com/elliot14A/tcorp/assessment/pkg/config"
	"github.com/elliot14A/tcorp/assessment/pkg/logger"
)

func StartClient() {
	configService := config.New()
	config, err := configService.LoadConfig()
	if err != nil {
		log.Fatal("error loading config file", "err", err.Error())
	}
	logger := logger.New()
	err = logger.SetConfig(config.GetLoggerConfig())
	if err != nil {
		log.Fatal("error setting logger config", "err", err.Error())
	}
	app := app.NewApp()
	app.SetConfig(config)
	app.SetLogger(logger)
	http.RunHttpServer(app)
}
