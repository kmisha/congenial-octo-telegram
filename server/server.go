package server

import (
	"auth/models"
	"auth/store"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

const (
	USER_PATH  = "/user"
	LOGIN_PATH = "/user/login"
)

type AuthServer struct {
	store     store.UsersStore
	tokenAuth jwtauth.JWTAuth
	Router    http.Handler
}

func NewAuthServer(store store.UsersStore, tokenAuth jwtauth.JWTAuth) *AuthServer {
	router := chi.NewRouter()
	s := &AuthServer{store, tokenAuth, router}

	router.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(&tokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Post(USER_PATH, s.createUser)
		r.Put(USER_PATH, s.updateUser)
	})

	router.Post(LOGIN_PATH, s.getJWTToken)

	return s
}

func (s *AuthServer) getJWTToken(w http.ResponseWriter, r *http.Request) {
	_, jwt, _ := s.tokenAuth.Encode(map[string]interface{}{"auth": true})
	w.Header().Set("jwt", jwt)
}

func (s *AuthServer) createUser(w http.ResponseWriter, r *http.Request) {
	var c CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.store.Create(c.Name, c.Login, c.Password, c.Phone, c.BirthDate)
}

func (s *AuthServer) updateUser(w http.ResponseWriter, r *http.Request) {
	var u UpdateUserRequest
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	patch := models.NewUser(u.Name, u.Login, u.Password, u.Phone, u.BirthDate)
	s.store.Update(u.ID, patch)
}
