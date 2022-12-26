package server

import (
	"auth/store"
	"io"
	"net/http"
	"strings"
)

const (
	PATH = "/user/"
)

type AuthServer struct {
	store store.UsersStore
}

func NewAuthServer(s store.UsersStore) *AuthServer {
	return &AuthServer{store: s}
}

func (s *AuthServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	isLoginPath := strings.HasPrefix(r.URL.Path, PATH+"login")
	isUserPath := strings.HasPrefix(r.URL.Path, PATH)

	if isLoginPath {
		s.processLogin(w, r)
		return
	}

	if isUserPath {
		s.processUser(w, r, getUserName(r))
	}

	w.WriteHeader(http.StatusBadGateway)
}

func getUserName(r *http.Request) string {
	return strings.TrimPrefix(r.URL.Path, PATH)
}

func (s *AuthServer) processUser(w http.ResponseWriter, r *http.Request, name string) {
	switch r.Method {
	case http.MethodPost:
		s.store.CreateUserRecord(name)
		w.WriteHeader(http.StatusCreated)
	case http.MethodPut:
		patch, _ := io.ReadAll(r.Body)
		s.store.UpdateUserRecord(name, string(patch))
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (s *AuthServer) processLogin(w http.ResponseWriter, r *http.Request) {}
