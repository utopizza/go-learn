package bloomfilter

import (
	"hash/fnv"
)

type BloomFilter struct {
	bitArray     []bool
	hashFuncList []HashFunc
}

type HashFunc func(s string) uint32

// 使用New32a实现第一个hash
func hash1(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

// 使用New32实现第一个hash
func hash2(s string) uint32 {
	h := fnv.New32()
	h.Write([]byte(s))
	return h.Sum32()
}

func NewBloomFilter(size uint32, hashFuncList []HashFunc) *BloomFilter {
	return &BloomFilter{
		bitArray:     make([]bool, size),
		hashFuncList: hashFuncList,
	}
}

// Add 添加一个元素
func (b *BloomFilter) Add(s string) {
	bitLen := uint32(len(b.bitArray))
	for _, hash := range b.hashFuncList {
		index := hash(s) % bitLen
		b.bitArray[index] = true
	}
}

// Exist 判断一个元素是否存在
func (b *BloomFilter) Exist(s string) bool {
	count := 0

	bitLen := uint32(len(b.bitArray))
	for _, hash := range b.hashFuncList {
		index := hash(s) % bitLen
		if b.bitArray[index] {
			count++
		}
	}

	if count == len(b.hashFuncList) { // 若每一个hash结果都命中，则判断为存在
		return true
	}

	return false
}
