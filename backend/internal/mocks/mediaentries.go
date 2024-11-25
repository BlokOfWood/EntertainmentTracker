package mocks

import (
	"github.com/BlokOfWood/EntertainmentTracker/backend/internal/data"
	"github.com/stretchr/testify/mock"
)

type MockMediaEntryModel struct {
	mock.Mock
}

func (m *MockMediaEntryModel) Insert(mediaEntry *data.MediaEntry) error {
	args := m.Called(mediaEntry)
	return args.Error(0)
}

func (m *MockMediaEntryModel) Get(id int64, userId int64) (*data.MediaEntry, error) {
	args := m.Called(id, userId)
	return args.Get(0).(*data.MediaEntry), args.Error(1)
}

func (m *MockMediaEntryModel) GetAll(userId int64) ([]*data.MediaEntry, error) {
	args := m.Called(userId)
	return args.Get(0).([]*data.MediaEntry), args.Error(1)
}

func (m *MockMediaEntryModel) Update(mediaEntry *data.MediaEntry) error {
	args := m.Called(mediaEntry)
	return args.Error(0)
}

func (m *MockMediaEntryModel) Delete(id int64, userId int64) error {
	args := m.Called(id, userId)
	return args.Error(0)
}
