package server

import (
	"log"

	"github.com/elliot14A/tcorp/assessment/pkg/config"
	"github.com/elliot14A/tcorp/assessment/pkg/logger"
	"github.com/elliot14A/tcorp/assessment/server/app"
	"github.com/elliot14A/tcorp/assessment/server/internal/infrastructure/database"
	"github.com/elliot14A/tcorp/assessment/server/internal/infrastructure/grpc"
)

func StartServer() {
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
	database := database.New()
	database.Connect(logger)
	app := app.NewApp()
	app.SetConfig(config)
	app.SetLogger(logger)
	app.SetDatabase(database)

	grpc.RunGrpcServer(app)
}
