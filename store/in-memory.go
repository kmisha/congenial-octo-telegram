package store

import (
	"auth/models"
	"sync"
	"time"

	"github.com/google/uuid"
)

type UsersStore interface {
	Create(name, login, pasword, phone string, birthDate time.Time) uuid.UUID
	Update(user, patch *models.User) uuid.UUID
}

type InMemoryUsersStore struct {
	mu    sync.Mutex
	store map[uuid.UUID]models.User
}

func NewInMemoryUsersStore() *InMemoryUsersStore {
	return &InMemoryUsersStore{
		sync.Mutex{},
		map[uuid.UUID]models.User{},
	}
}

func (s *InMemoryUsersStore) CreateUserRecord(name, login, password, phone string, birthDate time.Time) uuid.UUID {
	user := models.NewUser(name, login, password, phone, birthDate)
	s.manageChange(func() {
		s.store[user.ID] = *user
	})

	return user.ID
}

func (s *InMemoryUsersStore) UpdateUserRecord(id uuid.UUID, patch *models.User) uuid.UUID {
	s.manageChange(func() {
		user := s.store[id]
		user.Patch(patch)
		s.store[id] = user
	})

	return id
}

func (s *InMemoryUsersStore) manageChange(fn func()) {
	s.mu.Lock()
	defer s.mu.Unlock()

	fn()
}
