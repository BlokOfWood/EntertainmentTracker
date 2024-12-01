package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/BlokOfWood/EntertainmentTracker/backend/internal/data"
)

const (
	EMAIL    = "john.doe@gmail.com"
	PASSWORD = "password123"
	NAME     = "John Doe"
)

func newTestApplication(t *testing.T) *application {
	config := config{
		port: 8080,
		env:  "test",
		cors: struct{ trustedOrigins []string }{
			trustedOrigins: []string{"localhost:6969", "localhost:8080"},
		},
		auth: struct{ expireTime int }{
			expireTime: 60,
		},
	}
	db, err := openTestDB()
	if err != nil {
		t.Fatal(err)
	}
	return &application{
		//logger: slog.New(slog.NewTextHandler(io.Discard, nil)),
		logger: slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})),
		config: config,
		db:     db,
		models: data.NewModels(db),
	}
}

type testServer struct {
	*httptest.Server
}

func newTestServer(t *testing.T, h http.Handler) *testServer {
	ts := httptest.NewTLSServer(h)
	return &testServer{ts}
}

func (ts *testServer) get(t *testing.T, urlPath string) (int, http.Header, string) {
	rs, err := ts.Client().Get(ts.URL + urlPath)
	if err != nil {
		t.Fatal(err)
	}

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	body = bytes.TrimSpace(body)

	return rs.StatusCode, rs.Header, string(body)
}

func (ts *testServer) getWithAuth(t *testing.T, urlPath string, token string) (int, http.Header, string) {
	req, err := http.NewRequest("GET", ts.URL+urlPath, nil)
	if err != nil {
		t.Fatal(err)
	}
	bearer := "Bearer " + token
	req.Header.Set("Authorization", bearer)

	rs, err := ts.Client().Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	body = bytes.TrimSpace(body)

	return rs.StatusCode, rs.Header, string(body)
}

func (ts *testServer) post(t *testing.T, urlPath string, body io.Reader) (int, http.Header, string) {
	rs, err := ts.Client().Post(ts.URL+urlPath, "application/json", body)
	if err != nil {
		t.Fatal(err)
	}

	defer rs.Body.Close()
	bodyBytes, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	bodyBytes = bytes.TrimSpace(bodyBytes)

	return rs.StatusCode, rs.Header, string(bodyBytes)
}

func (ts *testServer) postWithAuth(t *testing.T, urlPath string, body io.Reader, token string) (int, http.Header, string) {
	req, err := http.NewRequest("POST", ts.URL+urlPath, body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+token)

	rs, err := ts.Client().Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer rs.Body.Close()
	bodyBytes, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	bodyBytes = bytes.TrimSpace(bodyBytes)

	return rs.StatusCode, rs.Header, string(bodyBytes)
}

func (ts *testServer) patchWithAuth(t *testing.T, urlPath string, body io.Reader, token string) (int, http.Header, string) {
	req, err := http.NewRequest("PATCH", ts.URL+urlPath, body)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer "+token)

	rs, err := ts.Client().Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer rs.Body.Close()
	bodyBytes, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	bodyBytes = bytes.TrimSpace(bodyBytes)

	return rs.StatusCode, rs.Header, string(bodyBytes)
}

func (ts *testServer) deleteWithAuth(t *testing.T, urlPath string, token string) (int, http.Header, string) {
	req, err := http.NewRequest("DELETE", ts.URL+urlPath, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+token)

	rs, err := ts.Client().Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer rs.Body.Close()
	bodyBytes, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	bodyBytes = bytes.TrimSpace(bodyBytes)

	return rs.StatusCode, rs.Header, string(bodyBytes)
}

func (ts *testServer) makeUser(t *testing.T, user *data.User, password string) (int, http.Header, string) {
	if user == nil {
		user = &data.User{
			Name:  NAME,
			Email: EMAIL,
		}
	}
	if password == "" {
		password = PASSWORD
	}

	input := map[string]interface{}{
		"name":     user.Name,
		"email":    user.Email,
		"password": password,
	}

	userJSON, err := json.Marshal(input)
	if err != nil {
		t.Fatal(err)
	}

	return ts.post(t, "/v1/users/register", bytes.NewReader(userJSON))
}

func (ts *testServer) loginUser(t *testing.T, email, password string) (int, http.Header, string) {
	if email == "" {
		email = EMAIL
	}
	if password == "" {
		password = PASSWORD
	}
	input := map[string]string{
		"email":    email,
		"password": password,
	}
	inputJSON, err := json.Marshal(input)
	if err != nil {
		t.Fatal(err)
	}

	return ts.post(t, "/v1/users/login", bytes.NewReader(inputJSON))
}

func (ts *testServer) getToken(t *testing.T, email, password string) string {
	if email == "" {
		email = EMAIL
	}

	if password == "" {
		password = PASSWORD
	}

	code, _, bodyString := ts.loginUser(t, email, password)
	if code != http.StatusCreated {
		t.Fatalf("unexpected status code %d", code)
	}

	var body map[string]interface{}
	err := json.Unmarshal([]byte(bodyString), &body)
	if err != nil {
		t.Fatal(err)
	}
	auth, ok := body["authentication_token"].(map[string]interface{})
	if !ok {
		t.Fatal("authentication_token not found in response")
	}

	token, ok := auth["token"].(string)
	if !ok {
		t.Fatal("token not found in authentication_token")
	}

	return token
}
