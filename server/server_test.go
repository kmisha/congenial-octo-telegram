package server

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d; want %d", got, want)
	}
}

func newPostUserRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/user/%s", name), nil)
	return request
}

func newPutUserRequest(name, newName string) *http.Request {
	buffer := bytes.NewBufferString(newName)
	request, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/user/%s", name), buffer)
	return request
}

type StubUsersStore struct {
	users map[string]string
}

func (s *StubUsersStore) CreateUserRecord(name string) {
	s.users[name] = name
}
func (s *StubUsersStore) UpdateUserRecord(name string, patch string) {
	s.users[name] = patch
}

func TestUserEndpoint(t *testing.T) {
	store := StubUsersStore{
		map[string]string{},
	}
	srv := NewAuthServer(&store)
	t.Run("returns a bad gateway error if path isn't login", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/qwerty", nil)
		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusBadGateway)
	})

	t.Run("returns the method not allowed error if using wrong method", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPatch, "/user/", nil)
		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusMethodNotAllowed)
	})

	t.Run("creates a user", func(t *testing.T) {
		name := "Bob"
		request := newPostUserRequest(name)
		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusCreated)
		got, ok := store.users[name]

		if !ok {
			t.Errorf("got %s; want %s", got, name)
		}
	})

	t.Run("updates a user", func(t *testing.T) {
		name := "Bob"
		want := "John"
		request := newPutUserRequest(name, want)
		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		got := store.users[name]

		if got != want {
			t.Errorf("got %s; want %s", got, want)
		}
	})
}
