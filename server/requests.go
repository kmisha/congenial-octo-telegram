package server

import (
	"time"

	"github.com/google/uuid"
)

type CreateUserRequest struct {
	Name      string    `json:"name"`
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
	BirthDate time.Time `json:"birthDate"`
}

type UpdateUserRequest struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
	BirthDate time.Time `json:"birthDate"`
}
