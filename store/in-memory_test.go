package store

import (
	"fmt"
	"sync"
	"testing"
)

func TestNewInMemoryUsersStore(t *testing.T) {
	got := NewInMemoryUsersStore()

	if got == nil {
		t.Error("fail to create a store")
	}
}

func TestCreateUserRecord(t *testing.T) {

	t.Run("we can add a user to the store", func(t *testing.T) {
		store := NewInMemoryUsersStore()
		user := "John"

		store.CreateUserRecord(user)

		_, ok := store.store[user]

		if !ok {
			t.Error("fail to create a user in store")
		}
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		store := NewInMemoryUsersStore()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(idx int) {
				defer wg.Done()
				name := fmt.Sprintf("John %d", idx+1)
				store.CreateUserRecord(name)
			}(i)
		}

		wg.Wait()

		if len(store.store) != wantedCount {
			t.Errorf("can't create %d records cuncurrently; got %d", wantedCount, len(store.store))
		}
	})
}

func TestUpdateUserRecord(t *testing.T) {

	t.Run("we can update a user in the store", func(t *testing.T) {
		store := NewInMemoryUsersStore()
		user := "John"
		want := "Bob"

		store.CreateUserRecord(user)
		store.UpdateUserRecord(user, want)

		got, ok := store.store[user]

		if !ok {
			t.Errorf("fail to read a user %s in store", user)
		}

		if got != want {
			t.Errorf("fail to read a user %s in store; got %s", want, got)
		}
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		store := NewInMemoryUsersStore()

		for i := 0; i < wantedCount; i++ {
			store.CreateUserRecord(fmt.Sprintf("John%d", i))
		}

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(idx int) {
				defer wg.Done()
				name := fmt.Sprintf("John%d", idx)
				patch := fmt.Sprintf("Bob%d", idx)
				store.UpdateUserRecord(name, patch)
			}(i)
		}

		wg.Wait()

		if len(store.store) != wantedCount {
			t.Errorf("can't create %d records cuncurrently; got %d", wantedCount, len(store.store))
		}
	})
}
