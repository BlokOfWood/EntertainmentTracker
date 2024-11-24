package mocks

import (
	"time"

	"github.com/BlokOfWood/EntertainmentTracker/backend/internal/data"
	"github.com/stretchr/testify/mock"
)

type MockTokenModel struct {
	mock.Mock
}

func (m *MockTokenModel) New(userID int64, ttl time.Duration) (*data.Token, error) {
	args := m.Called(userID, ttl)
	return args.Get(0).(*data.Token), args.Error(1)
}

func (m *MockTokenModel) Insert(token *data.Token) error {
	args := m.Called(token)
	return args.Error(0)
}

func (m *MockTokenModel) DeleteAllForUser(userID int64) error {
	args := m.Called(userID)
	return args.Error(0)
}
