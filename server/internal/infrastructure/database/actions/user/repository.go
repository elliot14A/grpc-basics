package user

import (
	"errors"

	"github.com/elliot14A/tcorp/assessment/server/internal/domain/models"
	"github.com/elliot14A/tcorp/assessment/server/internal/infrastructure/database"
)

type MockUserRepository struct {
	client []models.User
}

func NewMockUserRepository(db database.DatabasePort) (*MockUserRepository, error) {
	client, ok := db.GetClient().([]models.User)
	if !ok {
		return nil, errors.New("cannot cast database client to []models.User")
	}
	return &MockUserRepository{client: client}, nil
}
