package mocks

import (
	"time"

	"github.com/BlokOfWood/EntertainmentTracker/backend/internal/data"
	"github.com/stretchr/testify/mock"
)

var mockTime = time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

type MockUserModel struct {
	mock.Mock
}

func (m *MockUserModel) Insert(user *data.User) error {
	args := m.Called(user)

	if user != nil {
		user.ID = int64(1)
		user.CreatedAt = mockTime
		user.Version = 1
	}

	return args.Error(0)
}

func (m *MockUserModel) Get(id int64) (*data.User, error) {
	args := m.Called(id)
	return args.Get(0).(*data.User), args.Error(1)
}

func (m *MockUserModel) GetByEmail(email string) (*data.User, error) {
	args := m.Called(email)
	return args.Get(0).(*data.User), args.Error(1)
}

func (m *MockUserModel) Update(user *data.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserModel) GetForToken(tokenPlaintext string) (*data.User, error) {
	args := m.Called(tokenPlaintext)
	return args.Get(0).(*data.User), args.Error(1)
}
