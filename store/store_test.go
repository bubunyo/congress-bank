package store

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStore_InsertOk(t *testing.T) {
	s := NewStore[string]()

	ok := s.Insert("abcd", "efgh")
	assert.True(t, ok, "first insert failed")

	ok = s.Insert("abcd", "efgh")
	assert.False(t, ok, "second insert fialed")
}

func TestStore_InsertOkRace(t *testing.T) {
	s := NewStore[int]()

	count := 1000

	var wg sync.WaitGroup

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(i int) {
			assert.True(t, s.Insert(fmt.Sprintf("key_%d", i), i))
			wg.Done()
		}(i)
	}
	wg.Wait()
	assert.Equal(t, count, len(s.d))
}
