package store

import (
	"auth/models"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestNewInMemoryUsersStore(t *testing.T) {
	got := NewInMemoryUsersStore()

	if got == nil {
		t.Error("fail to create a store")
	}
}

func TestCreate(t *testing.T) {

	t.Run("we can add a user to the store", func(t *testing.T) {
		store := NewInMemoryUsersStore()
		name := "John"

		id := store.Create(name, "gwen", "123", "", time.Now())

		_, ok := store.store[id]

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
				store.Create(name, "gwen", "123", "", time.Now())
			}(i)
		}

		wg.Wait()

		if len(store.store) != wantedCount {
			t.Errorf("can't create %d records cuncurrently; got %d", wantedCount, len(store.store))
		}
	})
}

func TestUpdate(t *testing.T) {

	t.Run("we can update a user in the store", func(t *testing.T) {
		store := NewInMemoryUsersStore()
		user := "John"
		want := models.NewUser("Bob", "sten", "321", "123", time.Now())

		id := store.Create(user, "gwen", "123", "", time.Now())
		store.Update(id, want)

		got, ok := store.store[id]

		if !ok {
			t.Errorf("fail to read a user %s in store", user)
		}

		if got.Name != want.Name {
			t.Errorf("fail to read a user %s in store; got %s", want.Name, got.Name)
		}
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		store := NewInMemoryUsersStore()
		ids := make([]uuid.UUID, wantedCount)

		for i := 0; i < wantedCount; i++ {
			name := fmt.Sprintf("John%d", i)
			ids[i] = store.Create(name, "gwen", "123", "", time.Now())
		}

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for _, id := range ids {
			go func(id uuid.UUID) {
				defer wg.Done()

				patch := models.NewUser("Bob", "gwen", "123", "", time.Now())
				store.Update(id, patch)
			}(id)
		}

		wg.Wait()

		if len(store.store) != wantedCount {
			t.Errorf("can't create %d records cuncurrently; got %d", wantedCount, len(store.store))
		}
	})
}
