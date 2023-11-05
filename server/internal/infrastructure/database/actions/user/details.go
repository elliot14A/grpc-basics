package user

import (
	"errors"

	"github.com/elliot14A/tcorp/assessment/server/internal/domain/models"
)

func (m *MockUserRepository) Details(id uint) (*models.User, error) {
	for _, user := range m.client {
		if user.Id == id {
			return &user, nil
		}
	}

	return nil, errors.New("user not found")
}
