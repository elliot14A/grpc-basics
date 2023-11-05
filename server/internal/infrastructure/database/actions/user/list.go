package user

import "github.com/elliot14A/tcorp/assessment/server/internal/domain/models"

func (m *MockUserRepository) List() []models.User {
	return m.client
}
