package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/BlokOfWood/EntertainmentTracker/backend/internal/data"
	"github.com/stretchr/testify/assert"
)

var mockTime = time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

func TestHealthcheckHandler(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, bodyString := ts.get(t, "/v1/healthcheck")
	body := map[string]interface{}{}
	err := json.Unmarshal([]byte(bodyString), &body)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, code)
	assert.Equal(t, "available", body["status"])
	systemInfo, ok := body["system_info"].(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, "test", systemInfo["environment"])
	assert.Equal(t, "1.0.0", systemInfo["version"])
	corsTrustedOrigins, ok := systemInfo["cors_trusted_origins"].([]interface{})
	assert.True(t, ok)
	assert.Equal(t, "localhost:6969", corsTrustedOrigins[0])
	assert.Equal(t, "localhost:8080", corsTrustedOrigins[1])
}

func TestCreateUserHandler(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	input := map[string]string{
		"name":     "John Doe",
		"email":    "johndoe@gmail.com",
		"password": "password123",
	}
	inputJSON, err := json.Marshal(input)
	assert.NoError(t, err)

	code, _, bodyString := ts.post(t, "/v1/users/register", bytes.NewReader(inputJSON))

	assert.Equal(t, http.StatusCreated, code)

	var body map[string]interface{}
	err = json.Unmarshal([]byte(bodyString), &body)
	assert.NoError(t, err)
	user, ok := body["user"].(map[string]interface{})
	assert.True(t, ok)

	assert.NotEmpty(t, user["created_at"])
	assert.Equal(t, float64(1), user["id"])
	assert.Equal(t, "johndoe@gmail.com", user["email"])
	assert.Equal(t, "John Doe", user["name"])
}

func TestCreateUserHandler_InvalidJSON(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, _ := ts.post(t, "/v1/users/register", bytes.NewReader([]byte("invalid json")))
	assert.Equal(t, http.StatusBadRequest, code)
}

func TestCreateUserHandler_ValidationError(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	input := map[string]string{
		"name":     "",
		"email":    "invalidemail",
		"password": "123",
	}
	inputJSON, err := json.Marshal(input)
	assert.NoError(t, err)

	code, _, _ := ts.post(t, "/v1/users/register", bytes.NewReader(inputJSON))
	assert.Equal(t, http.StatusUnprocessableEntity, code)
}

func TestCreateUserHandler_DuplicateEmail(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	input := map[string]string{
		"name":     "John Doe",
		"email":    "johndoe@example.com",
		"password": "securepassword",
	}
	inputJSON, err := json.Marshal(input)
	assert.NoError(t, err)

	code, _, _ := ts.post(t, "/v1/users/register", bytes.NewReader(inputJSON))
	assert.Equal(t, http.StatusCreated, code)

	code, _, _ = ts.post(t, "/v1/users/register", bytes.NewReader(inputJSON))
	assert.Equal(t, http.StatusConflict, code)
}

func TestCreateUserHandler_DBError(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	app.db.Close()

	input := map[string]string{
		"name":     "John Doe",
		"email":    "john.doe@gmail.com",
		"password": "password123",
	}

	inputJSON, err := json.Marshal(input)
	assert.NoError(t, err)

	code, _, _ := ts.post(t, "/v1/users/register", bytes.NewReader(inputJSON))
	assert.Equal(t, http.StatusInternalServerError, code)
}

func TestCreateAuthenticationTokenHandler(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, _ := ts.makeUser(t, &data.User{
		Email: "john.doe@gmail.com",
		Name:  "John Doe",
	}, "password123")
	assert.Equal(t, http.StatusCreated, code)

	input := map[string]string{
		"email":    "john.doe@gmail.com",
		"password": "password123",
	}
	inputJSON, err := json.Marshal(input)
	assert.NoError(t, err)

	code, _, bodyString := ts.post(t, "/v1/users/login", bytes.NewReader(inputJSON))
	assert.Equal(t, http.StatusCreated, code)

	var body map[string]interface{}
	err = json.Unmarshal([]byte(bodyString), &body)
	assert.NoError(t, err)
	auth, ok := body["authentication_token"].(map[string]interface{})
	assert.True(t, ok)

	assert.NotEmpty(t, auth["token"])
	assert.NotEmpty(t, auth["expiry"])
}

func TestCreateAuthenticationTokenHandler_InvalidJSON(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, _ := ts.post(t, "/v1/users/login", bytes.NewReader([]byte("invalid json")))
	assert.Equal(t, http.StatusBadRequest, code)
}

func TestCreateAuthenticationTokenHandler_ValidationError(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	input := map[string]string{
		"email":    "invalidemail",
		"password": "123",
	}
	inputJSON, err := json.Marshal(input)
	assert.NoError(t, err)

	code, _, _ := ts.post(t, "/v1/users/login", bytes.NewReader(inputJSON))
	assert.Equal(t, http.StatusUnprocessableEntity, code)
}

func TestCreateAuthenticationTokenHandler_UserNotFound(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	input := map[string]string{
		"email":    "joe.biden@gmail.com",
		"password": "password123",
	}
	inputJSON, err := json.Marshal(input)
	assert.NoError(t, err)

	code, _, _ := ts.post(t, "/v1/users/login", bytes.NewReader(inputJSON))
	assert.Equal(t, http.StatusUnauthorized, code)
}

func TestCreateAuthenticationTokenHandler_IncorrectPassword(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, _ := ts.makeUser(t, nil, "")
	assert.Equal(t, http.StatusCreated, code)

	input := map[string]string{
		"email":    "john.doe@gmail.com",
		"password": "incorrectpassword",
	}
	inputJSON, err := json.Marshal(input)
	assert.NoError(t, err)

	code, _, _ = ts.post(t, "/v1/users/login", bytes.NewReader(inputJSON))
	assert.Equal(t, http.StatusUnauthorized, code)
}

func TestCreateAuthenticationTokenHandler_DBError(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	app.db.Close()

	input := map[string]string{
		"email":    "john.doe@gmail.com",
		"password": "password123",
	}
	inputJSON, err := json.Marshal(input)
	assert.NoError(t, err)

	code, _, _ := ts.post(t, "/v1/users/login", bytes.NewReader(inputJSON))
	assert.Equal(t, http.StatusInternalServerError, code)
}

func TestGetUserHandler(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, _ := ts.makeUser(t, nil, "")
	assert.Equal(t, http.StatusCreated, code)

	token := ts.getToken(t, "", "")

	code, _, bodyString := ts.getWithAuth(t, "/v1/users/me", token)

	var body map[string]interface{}
	err := json.Unmarshal([]byte(bodyString), &body)
	assert.NoError(t, err)
	user, ok := body["user"].(map[string]interface{})
	assert.True(t, ok)

	assert.Equal(t, "John Doe", user["name"])
	assert.Equal(t, "john.doe@gmail.com", user["email"])
	assert.Equal(t, float64(1), user["id"])
	assert.NotEmpty(t, user["created_at"])
}

func TestGetUserHandler_Unauthenticated(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, _ := ts.get(t, "/v1/users/me")
	assert.Equal(t, http.StatusUnauthorized, code)
}

func TestGetUserHandler_DBError(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, _ := ts.makeUser(t, nil, "")
	assert.Equal(t, http.StatusCreated, code)

	token := ts.getToken(t, "", "")

	app.db.Close()

	code, _, _ = ts.getWithAuth(t, "/v1/users/me", token)
	assert.Equal(t, http.StatusInternalServerError, code)
}

func TestInvalidateAuthenticationTokenHandler(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, _ := ts.makeUser(t, nil, "")
	assert.Equal(t, http.StatusCreated, code)

	token := ts.getToken(t, "", "")

	code, _, _ = ts.getWithAuth(t, "/v1/users/logout", token)
	assert.Equal(t, http.StatusOK, code)

	code, _, _ = ts.getWithAuth(t, "/v1/users/me", token)
	assert.Equal(t, http.StatusUnauthorized, code)
}

func TestInvalidateAuthenticationTokenHandler_Unauthenticated(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, _ := ts.get(t, "/v1/users/logout")
	assert.Equal(t, http.StatusUnauthorized, code)
}

func TestInvalidateAuthenticationTokenHandler_DBError(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, _ := ts.makeUser(t, nil, "")
	assert.Equal(t, http.StatusCreated, code)

	token := ts.getToken(t, "", "")

	app.db.Close()

	code, _, _ = ts.getWithAuth(t, "/v1/users/logout", token)
	assert.Equal(t, http.StatusInternalServerError, code)
}
