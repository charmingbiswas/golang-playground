package lru

import (
	"reflect"
	"testing"
)

// should return a valid pointer to LRUCache struct, with capacity 10
func TestLRUCache(t *testing.T) {
	capacity := 10
	lru := Constructor(capacity)

	if lru.capacity != 10 {
		t.Errorf("LRU capacity not equal to %d", capacity)
	}

	if reflect.TypeOf(lru) != reflect.TypeFor[*LRUCache]() {
		t.Error("LRU cache type mismatch")
	}
}
