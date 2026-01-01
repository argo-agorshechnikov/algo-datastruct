package main

import (
	"fmt"
	"hash/fnv"
)

// key-value struct
type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

// Chain of key-value pairs
type Bucket[K comparable, V any] struct {
	pairs []*Pair[K, V]
}

type HashTable[K comparable, V any] struct {
	buckets []*Bucket[K, V]
	count   int
	size    int
}

// Create hash table with started size
func NewHashTable[K comparable, V any](initSize int) *HashTable[K, V] {
	size := nextPrime(initSize)
	buckets := make([]*Bucket[K, V], size)
	for i := range buckets {
		buckets[i] = &Bucket[K, V]{}
	}

	return &HashTable[K, V]{buckets: buckets, size: size}
}

// Next simple int
func nextPrime(n int) int {
	if n < 2 {
		return 2
	}

	for {
		ok := true
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				ok = false
				break
			}
		}

		if ok {
			return n
		}
		n++
	}
}

// Hash function get hash-code
func (ht *HashTable[K, V]) hash(key K) uint32 {
	h := fnv.New32a()
	h.Write([]byte(fmt.Sprintf("%v", key)))
	return h.Sum32() % uint32(ht.size)
}

// Put - insert or update value
func (ht *HashTable[K, V]) Put(key K, value V) {
	idx := int(ht.hash(key))
	bucket := ht.buckets[idx]

	// Check current
	for _, p := range bucket.pairs {
		if p.Key == key {
			p.Value = value
			return
		}
	}

	// Add new
	bucket.pairs = append(bucket.pairs, &Pair[K, V]{Key: key, Value: value})
	ht.count++

	// Resize
	if float64(ht.count)/float64(ht.size) > 0.7 {
		ht.resize()
	}

}

func (ht *HashTable[K, V]) resize() {
	oldBuckets := ht.buckets
	ht.size = nextPrime(ht.size * 2)
	ht.buckets = make([]*Bucket[K, V], ht.size)

	for i := range ht.buckets {
		ht.buckets[i] = &Bucket[K, V]{}
	}

	ht.count = 0

	for _, bucket := range oldBuckets {
		for _, p := range bucket.pairs {
			ht.Put(p.Key, p.Value)
		}
	}
}

func (ht *HashTable[K, V]) Len() int {
	return ht.count
}

// Get value or return panic
func (ht *HashTable[K, V]) Get(key K) V {
	idx := int(ht.hash(key))

	for _, p := range ht.buckets[idx].pairs {
		if p.Key == key {
			return p.Value
		}
	}

	panic("Key not found")
}

// Delete key
func (ht *HashTable[K, V]) Delete(key K) {
	idx := int(ht.hash(key))
	bucket := ht.buckets[idx]

	for i, p := range bucket.pairs {
		if p.Key == key {
			bucket.pairs = append(bucket.pairs[:i], bucket.pairs[i+1:]...)
			ht.count--
			return
		}
	}
}

func (p *Pair[K, V]) String() string {
	return fmt.Sprintf("(%v:%v)", p.Key, p.Value)
}

func (ht *HashTable[K, V]) Print() {
	fmt.Printf("HashTable: %d/%d элементов (load=%.2f)\n", ht.count, ht.size, float64(ht.count)/float64(ht.size))
	for i, bucket := range ht.buckets {
		if len(bucket.pairs) > 0 {
			fmt.Printf("Бакет %2d: %v\n", i, bucket.pairs)
		}
	}
}

func main() {
	ht := NewHashTable[string, int](4)

	ht.Put("one", 1)
	ht.Put("two", 2)
	ht.Print()

	fmt.Println("--------------------------------")
	fmt.Println("Update table")
	ht.Put("one", 99)
	ht.Print()

	fmt.Println("----------------------------------")
	ht.Put("four", 4)
	fmt.Println("Len after Put: ", ht.Len())

	fmt.Println("--------------------------------")
	fmt.Println("Get two:", ht.Get("two"))

	fmt.Println("--------------------------------")
	fmt.Println("Delete 'one' key")
	ht.Delete("one")
	ht.Print()

	fmt.Println("--------------------------------")
	fmt.Println("Resize test")
	for i := 0; i < 20; i++ {
        ht.Put(fmt.Sprintf("key%d", i), i)
    }

	fmt.Println("Size: ", ht.size)
	fmt.Println("Len: ", ht.Len())
	ht.Print()
}
