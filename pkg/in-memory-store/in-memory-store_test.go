package inmemorystore

import (
	"fmt"
	"sync"
	"testing"
)

func TestNewInMemoryStore(t *testing.T) {
	store := NewInMemoryUsersStore[string, string]()
	key := "Bob"
	value := "Yes"
	var got string

	store.Update(func(tx *Tx[string, string]) error {
		tx.Set(key, value)
		return nil
	})

	store.View(func(tx *Tx[string, string]) error {
		got = tx.Get(key)
		return nil
	})

	if got != value {
		t.Errorf(`can't read correct value; want "%q" but got "%s"`, value, got)
	}
}

func TestInMemoryStore(t *testing.T) {
	t.Run("it runs updates safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		store := NewInMemoryUsersStore[string, int]()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(idx int) {
				defer wg.Done()
				name := fmt.Sprintf("John%d", idx)
				store.Update(func(tx *Tx[string, int]) error {
					tx.Set(name, idx)
					return nil
				})
			}(i)
		}

		wg.Wait()

		if len(store.store) != wantedCount {
			t.Errorf("can't create %d records cuncurrently; got %d", wantedCount, len(store.store))
		}
	})

	t.Run("it runs views and updates safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		store := NewInMemoryUsersStore[string, int]()
		mu := sync.Mutex{}
		expected := make(map[string]int)

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(idx int) {
				defer wg.Done()
				name := fmt.Sprintf("John%d", idx)
				store.Update(func(tx *Tx[string, int]) error {
					tx.Set(name, idx)
					tx.Get(name)
					return nil
				})

				var value int
				store.View(func(tx *Tx[string, int]) error {
					value = tx.Get(name)
					return nil
				})

				mu.Lock()
				expected[name] = value
				mu.Unlock()
			}(i)
		}

		wg.Wait()

		if len(store.store) != wantedCount {
			t.Errorf("can't create %d records cuncurrently; got %d", wantedCount, len(store.store))
		}

		for k, v := range expected {
			if v != store.store[k] {
				t.Errorf("found a deviation in store; want %d but got %d", v, store.store[k])
			}
		}
	})
}
