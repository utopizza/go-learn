package lru

import "testing"

func TestNewCache(t *testing.T) {
	c := NewCache(3)
	c.Put("1", 1)
	c.Put("2", 2)
	c.Put("3", 3)
	c.Get("2")
	c.Put("4", 4)
	c.Put("3", 3)
}
