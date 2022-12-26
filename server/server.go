package server

import (
	"auth/store"
	"io"
	"net/http"
	"strings"
)

type AuthServer struct {
	store store.UsersStore
}

func NewAuthServer(s store.UsersStore) *AuthServer {
	return &AuthServer{store: s}
}

func (s *AuthServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	isUserPath := strings.HasPrefix(r.URL.Path, "/user")

	if !isUserPath {
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	name := strings.TrimPrefix(r.URL.Path, "/user/")

	switch r.Method {
	case http.MethodPost:
		s.store.CreateUserRecord(name)
	case http.MethodPut:
		patch, _ := io.ReadAll(r.Body)
		s.store.UpdateUserRecord(name, string(patch))
	}
}
