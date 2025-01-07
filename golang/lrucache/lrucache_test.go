package lrucache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCaching(t *testing.T) {
	cache := New[int](3)

	cache.Put("a", 100)
	cache.Put("b", 105)
	cache.Put("c", 110)
	v, ok := cache.Get("a")
	assert.Equal(t, true, ok, "should have found a")
	assert.Equal(t, 100, v, "should have stored the correct value for a")

	v, ok = cache.Get("b")
	assert.Equal(t, true, ok, "should have found b")
	assert.Equal(t, 105, v, "should have stored the correct value for b ")

	v, ok = cache.Get("c")
	assert.Equal(t, true, ok, "should have found c")
	assert.Equal(t, 110, v, "should have stored the correct value for c")

	cache.Put("d", 115)
	_, ok = cache.Get("a")
	assert.Equal(t, false, ok, "a should have been GC'd")

	v, ok = cache.Get("d")
	assert.Equal(t, true, ok, "should have found d")
	assert.Equal(t, 115, v, "should have stored the correct value for d")

	cache.Del("b")
	_, ok = cache.Get("b")
	assert.Equal(t, false, ok, "b should have been removed")
}
