package main

import "fmt"

// TODO: Define Cache interface, MapCache struct, NewMapCache constructor, and implement all methods
type Cache[K comparable, V any] interface {
	Get(key K) (V, bool)
	Set(key K, val V)
	Delete(key K)
	Len() int
}

type MapCache[K comparable, V any] struct {
	CacheMap map[K]V
}

func NewMapCache[K comparable, V any]() *MapCache[K, V] {
	NewMap := MapCache[K, V]{
		CacheMap: make(map[K]V),
	}
	return &NewMap
}

func (m *MapCache[K, V]) Get(key K) (V, bool) {
	val, ok := m.CacheMap[key]
	if ok {
		return val, true
	}
	return val, false
}

func (m *MapCache[K, V]) Set(key K, val V) {
	m.CacheMap[key] = val
}

func (m *MapCache[K, V]) Delete(key K) {
	delete(m.CacheMap, key)
}

func (m *MapCache[K, V]) Len() int {
	count := 0
	for range m.CacheMap {
		count++
	}
	return count
}

func main() {
	// Use interface type to prove MapCache implements Cache
	var cache Cache[string, string] = NewMapCache[string, string]()

	fmt.Println("Set name=Alice, age=30")
	cache.Set("name", "Alice")
	cache.Set("age", "30")

	val, ok := cache.Get("name")
	fmt.Printf("Get name: %s, %v\n", val, ok)

	val, ok = cache.Get("missing")
	fmt.Printf("Get missing: %s, %v\n", val, ok)

	fmt.Printf("Len: %d\n", cache.Len())

	fmt.Println("Delete age")
	cache.Delete("age")

	fmt.Printf("Len: %d\n", cache.Len())

	// Different types
	intCache := NewMapCache[string, int]()
	intCache.Set("count", 100)
	n, _ := intCache.Get("count")
	fmt.Printf("Int cache: %d, true\n", n)
}
