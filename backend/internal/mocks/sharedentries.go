package mocks

import (
	"github.com/BlokOfWood/EntertainmentTracker/backend/internal/data"
	"github.com/stretchr/testify/mock"
)

type MockSharedEntryModel struct {
	mock.Mock
}

func (m *MockSharedEntryModel) Insert(entryID, sharedBy, sharedWith int64) error {
	args := m.Called(entryID, sharedBy, sharedWith)
	return args.Error(0)
}

func (m *MockSharedEntryModel) Get(id int64) (*data.SharedEntry, error) {
	args := m.Called(id)
	return args.Get(0).(*data.SharedEntry), args.Error(1)
}

func (m *MockSharedEntryModel) GetBySharedWithID(sharedWith int64) ([]*data.SharedEntry, error) {
	args := m.Called(sharedWith)
	return args.Get(0).([]*data.SharedEntry), args.Error(1)
}

func (m *MockSharedEntryModel) GetAll() ([]*data.SharedEntry, error) {
	args := m.Called()
	return args.Get(0).([]*data.SharedEntry), args.Error(1)
}

func (m *MockSharedEntryModel) Update(sharedEntry *data.SharedEntry) error {
	args := m.Called(sharedEntry)
	return args.Error(0)
}

func (m *MockSharedEntryModel) Delete(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}
