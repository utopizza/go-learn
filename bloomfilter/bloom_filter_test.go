package bloomfilter

import "testing"

func TestNewBloomFilter(t *testing.T) {

	f := NewBloomFilter(100, []HashFunc{hash1, hash2})
	f.Add("hello")
	f.Add("good bye")

	if !f.Exist("hello") {
		t.Error()
	}
	if !f.Exist("good bye") {
		t.Error()
	}
	if f.Exist("hello?") {
		t.Error()
	}
}
