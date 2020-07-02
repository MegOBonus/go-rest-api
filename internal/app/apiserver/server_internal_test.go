package apiserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/MegOBonus/go-rest-api/internal/app/model"
	"github.com/MegOBonus/go-rest-api/internal/app/store/teststore"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_AuthenticateUser(t *testing.T) {
	u := model.TesUser(t)
	store := teststore.New()
	store.User().Create(u)

	secretKey := []byte("secret")
	s := newServer(store, sessions.NewCookieStore(secretKey))
	sc := securecookie.New(secretKey, nil)
	handler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	})

	testCases := []struct {
		name        string
		cookieValue map[interface{}]interface{}
		code        int
	}{
		{
			name: "authenticated",
			cookieValue: map[interface{}]interface{}{
				"user_id": u.ID,
			},
			code: http.StatusOK,
		},
		{
			name:        "not authenticated",
			cookieValue: nil,
			code:        http.StatusUnauthorized,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/", nil)

			cookieStr, _ := sc.Encode(sessionName, test.cookieValue)
			req.Header.Set("Cookie", fmt.Sprintf("%s=%s", sessionName, cookieStr))
			s.authenticateUser(handler).ServeHTTP(rec, req)
			assert.Equal(t, test.code, rec.Code)
		})
	}
}

func TestServer_HandleUsersCreate(t *testing.T) {
	s := newServer(teststore.New(), sessions.NewCookieStore([]byte("Secret")))

	testCases := []struct {
		name    string
		payload interface{}
		code    int
	}{
		{
			name: "Valid",
			payload: map[string]string{
				"email":    "test@test.com",
				"password": "passw1234",
			},
			code: http.StatusCreated,
		}, {
			name:    "Invalid payload",
			payload: nil,
			code:    http.StatusUnprocessableEntity,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(test.payload)
			req, _ := http.NewRequest(http.MethodPost, "/users", b)

			s.ServeHTTP(rec, req)
			assert.Equal(t, test.code, rec.Code)
		})
	}
}

func TestServer_HandleSessionsCreate(t *testing.T) {
	u := model.TesUser(t)
	store := teststore.New()
	store.User().Create(u)
	s := newServer(store, sessions.NewCookieStore([]byte("Secret")))

	testCases := []struct {
		name    string
		payload interface{}
		code    int
	}{
		{
			name: "Valid",
			payload: map[string]string{
				"email":    u.Email,
				"password": u.Password,
			},
			code: http.StatusOK,
		}, {
			name:    "Invalid payload",
			payload: nil,
			code:    http.StatusUnauthorized,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(test.payload)
			req, _ := http.NewRequest(http.MethodPost, "/sessions", b)

			s.ServeHTTP(rec, req)
			assert.Equal(t, test.code, rec.Code)
		})
	}
}
