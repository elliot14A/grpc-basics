package app

import (
	"github.com/elliot14A/tcorp/assessment/pkg/config"
	"github.com/elliot14A/tcorp/assessment/pkg/logger"
	"github.com/elliot14A/tcorp/assessment/server/internal/infrastructure/database"
)

type App struct {
	logger   logger.LoggerPort
	config   config.ConfigPort
	database database.DatabasePort
}

func NewApp() *App {
	return &App{}
}

func (a *App) SetLogger(logger logger.LoggerPort) {
	a.logger = logger
}

func (a *App) SetConfig(config config.ConfigPort) {
	a.config = config
}

func (a *App) SetDatabase(database database.DatabasePort) {
	a.database = database
}

func (a *App) GetLogger() logger.LoggerPort {
	return a.logger
}

func (a *App) GetConfig() config.ConfigPort {
	return a.config
}

func (a *App) GetDatabase() database.DatabasePort {
	return a.database
}
