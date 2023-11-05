package database

import (
	"github.com/elliot14A/tcorp/assessment/pkg/logger"
	"github.com/elliot14A/tcorp/assessment/server/internal/domain/models"
)

type DatabasePort interface {
	Connect(logger logger.LoggerPort) error
	GetClient() interface{}
}

type MockDatabaseAdapter struct {
	Client []models.User
}

func (m *MockDatabaseAdapter) Connect(logger logger.LoggerPort) error {

	logger.Info("mocking connection to database")
	m.Client = append(m.Client, models.User{
		Id:        1,
		FirstName: "Elliot",
		City:      "Hyderabad",
		Height:    5.5,
		Married:   false,
		Phone:     "+919191919919",
	}, models.User{
		Id:        2,
		FirstName: "John",
		City:      "Newyork",
		Height:    6.1,
		Married:   true,
		Phone:     "+449191919919",
	}, models.User{
		Id:        3,
		FirstName: "Andrew",
		City:      "Humburg",
		Height:    5.9,
		Married:   false,
		Phone:     "+299191919919",
	}, models.User{
		Id:        4,
		FirstName: "Rapole",
		City:      "Hyderabad",
		Height:    5.8,
		Married:   false,
		Phone:     "+19191919919",
	})
	logger.Info("successfully mocked connection to database....")
	return nil
}

func (m *MockDatabaseAdapter) GetClient() interface{} {
	return m.Client
}

func New() DatabasePort {
	return &MockDatabaseAdapter{}
}
