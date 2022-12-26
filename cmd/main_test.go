package main

import (
	"auth/server"
	"auth/store"
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

// TODO: move to utils
func newPostUserRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/user/%s", name), nil)
	return request
}

// TODO: move to utils
func newPutUserRequest(name, newName string) *http.Request {
	buffer := bytes.NewBufferString(newName)
	request, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/user/%s", name), buffer)
	return request
}

// TODO: move to utils
func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d; want %d", got, want)
	}
}

func TestWorkWithServer(t *testing.T) {
	str := store.NewInMemoryUsersStore()
	srv := server.NewAuthServer(str)
	var wg sync.WaitGroup
	names := []string{"John", "Bob", "Alice"}
	var responses []*httptest.ResponseRecorder
	wg.Add(len(names))

	// do some staff
	for i, n := range names {
		go func(idx int, name string) {
			defer wg.Done()
			res := httptest.NewRecorder()
			responses = append(responses, res)
			srv.ServeHTTP(res, newPostUserRequest(name))
		}(i, n)
	}

	// check statuses
	for _, r := range responses {
		assertStatus(t, r.Code, http.StatusCreated)
	}
}
