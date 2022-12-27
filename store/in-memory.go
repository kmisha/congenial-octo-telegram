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
	s.manageChange(func() {
		s.store[name] = name
	})
}

func (s *InMemoryUsersStore) UpdateUserRecord(name string, patch string) {
	s.manageChange(func() {
		s.store[name] = patch
	})
}

func (s *InMemoryUsersStore) manageChange(fn func()) {
	s.mu.Lock()
	defer s.mu.Unlock()

	fn()
}
