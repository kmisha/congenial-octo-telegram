package inmemorystore

import "sync"

// Transaction
type Tx[K comparable, V any] struct {
	s        *InMemoryStore[K, V]
	writable bool
}

func (tx *Tx[K, V]) Set(key K, value V) {
	tx.s.store[key] = value
}

func (tx *Tx[K, V]) Get(key K) V {
	return tx.s.store[key]
}

func (tx *Tx[K, V]) lock() {
	if tx.writable {
		tx.s.mu.Lock()
	} else {
		tx.s.mu.RLock()
	}
}

func (tx *Tx[K, V]) unlock() {
	if tx.writable {
		tx.s.mu.Unlock()
	} else {
		tx.s.mu.RUnlock()
	}
}

// Store
type InMemoryStore[K comparable, V any] struct {
	mu    sync.RWMutex
	store map[K]V
}

func NewInMemoryUsersStore[K comparable, V any]() *InMemoryStore[K, V] {
	return &InMemoryStore[K, V]{sync.RWMutex{}, map[K]V{}}
}

func (s *InMemoryStore[K, V]) manageTransaction(writeble bool, fn func(tx *Tx[K, V]) error) (err error) {
	tx := s.startTransaction(writeble)
	defer tx.unlock()

	return fn(tx)
}

func (s *InMemoryStore[K, V]) startTransaction(writable bool) *Tx[K, V] {
	tx := &Tx[K, V]{s, writable}
	tx.lock()
	return tx
}

func (s *InMemoryStore[K, V]) View(fn func(tx *Tx[K, V]) error) error {
	return s.manageTransaction(false, fn)
}

func (s *InMemoryStore[K, V]) Update(fn func(tx *Tx[K, V]) error) error {
	return s.manageTransaction(true, fn)
}
