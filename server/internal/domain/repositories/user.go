package repositories

import (
	"github.com/elliot14A/tcorp/assessment/server/internal/domain/models"
	"github.com/elliot14A/tcorp/assessment/server/internal/infrastructure/database"
	"github.com/elliot14A/tcorp/assessment/server/internal/infrastructure/database/actions/user"
)

type UserRepository interface {
	Details(id uint) (*models.User, error)
	List() []models.User
}

func NewUserRepository(db database.DatabasePort) (UserRepository, error) {
	return user.NewMockUserRepository(db)
}
