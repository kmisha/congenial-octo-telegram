package store

import "sync"

type UsersStore interface {
	CreateUserRecord(name string)
	UpdateUserRecord(name, patch string)
}

type InMemoryUsersStore struct {
	mu    sync.Mutex
	store map[string]string
}

func NewInMemoryUsersStore() *InMemoryUsersStore {
	return &InMemoryUsersStore{
		sync.Mutex{},
		map[string]string{},
	}
}

func (s *InMemoryUsersStore) CreateUserRecord(name string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.store[name] = name
}

func (s *InMemoryUsersStore) UpdateUserRecord(name string, patch string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.store[name] = patch
}
