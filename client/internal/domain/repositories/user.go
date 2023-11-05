package repositories

import (
	"github.com/elliot14A/tcorp/assessment/client/internal/domain/models"
	"github.com/elliot14A/tcorp/assessment/client/internal/infrastructure/grpc"
	"github.com/elliot14A/tcorp/assessment/pkg/config"
	"github.com/elliot14A/tcorp/assessment/pkg/logger"
)

type UserRepository interface {
	Details(userId int) (*models.User, error)
	List() ([]models.User, error)
}

func NewUserRepository(config config.ServerConfig, logger logger.LoggerPort) UserRepository {
	repo := grpc.NewGrpcUserRepository(config, logger)
	return repo
}
