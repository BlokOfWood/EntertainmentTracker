package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/BlokOfWood/EntertainmentTracker/backend/internal/data"
	"github.com/BlokOfWood/EntertainmentTracker/backend/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var mockTime = time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

func TestHealthcheckHandler(t *testing.T) {
	app := &application{
		config: config{
			env: "development",
			cors: struct {
				trustedOrigins []string
			}{
				trustedOrigins: []string{"http://example.com"},
			},
		},
	}
	req := httptest.NewRequest(http.MethodGet, "/healthcheck", nil)
	w := httptest.NewRecorder()

	app.healthcheckHandler(w, req)

	res := w.Result()
	require.Equal(t, http.StatusOK, res.StatusCode)

	var body map[string]interface{}
	err := json.NewDecoder(res.Body).Decode(&body)
	require.NoError(t, err)
	assert.Equal(t, "available", body["status"])
	systemInfo := body["system_info"].(map[string]interface{})
	assert.Equal(t, "development", systemInfo["environment"])
	assert.Contains(t, systemInfo["cors_trusted_origins"], "http://example.com")
}

func TestCreateUserHandler(t *testing.T) {
	mockUserModel := new(mocks.MockUserModel)

	app := &application{
		models: data.Models{
			Users: mockUserModel,
		},
	}

	input := map[string]string{
		"name":     "John Doe",
		"email":    "johndoe@example.com",
		"password": "password123",
	}
	body, _ := json.Marshal(input)

	req := httptest.NewRequest(http.MethodPost, "/users/register", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	mockUserModel.On("Insert", mock.MatchedBy(func(user *data.User) bool {
		okPassword, _ := user.Password.Matches("password123")
		return user.Name == "John Doe" && user.Email == "johndoe@example.com" && okPassword
	})).Return(nil)

	app.createUserHandler(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var actualResponse map[string]interface{}
	err := json.Unmarshal(rr.Body.Bytes(), &actualResponse)
	assert.NoError(t, err)
	user, ok := actualResponse["user"].(map[string]interface{})
	assert.True(t, ok)

	assert.Equal(t, mockTime.Format(time.RFC3339), user["created_at"])
	assert.Equal(t, float64(1), user["id"])
	assert.Equal(t, "johndoe@example.com", user["email"])
	assert.Equal(t, "John Doe", user["name"])

	mockUserModel.AssertCalled(t, "Insert", mock.MatchedBy(func(user *data.User) bool {
		return user.Name == "John Doe" && user.Email == "johndoe@example.com"
	}))
}

func TestCreateUserHandler_InvalidJSON(t *testing.T) {
	app := &application{
		models: data.Models{},
		logger: slog.New(slog.NewTextHandler(io.Discard, nil)),
	}

	req := httptest.NewRequest(http.MethodPost, "/users/register", bytes.NewReader([]byte("{invalid json")))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	app.createUserHandler(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestCreateUserHandler_ValidationError(t *testing.T) {
	app := &application{
		models: data.Models{},
	}

	input := map[string]string{
		"name":     "",
		"email":    "invalidemail",
		"password": "123",
	}
	body, _ := json.Marshal(input)

	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	app.createUserHandler(rr, req)

	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
}

func TestCreateUserHandler_DuplicateEmail(t *testing.T) {
	mockUserModel := new(mocks.MockUserModel)
	mockUserModel.On("Insert", mock.Anything).Return(data.ErrDuplicateEmail)

	app := &application{
		models: data.Models{
			Users: mockUserModel,
		},
	}

	input := map[string]string{
		"name":     "John Doe",
		"email":    "johndoe@example.com",
		"password": "securepassword",
	}
	body, _ := json.Marshal(input)

	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	app.createUserHandler(rr, req)

	assert.Equal(t, http.StatusConflict, rr.Code)
	mockUserModel.AssertCalled(t, "Insert", mock.Anything)
}
