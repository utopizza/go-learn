package design

import "testing"

func Test(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Get(1)
	cache.Put(3, 3)
	cache.Get(2)
}
