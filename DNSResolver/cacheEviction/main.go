package main

import "fmt"

type Record struct {
	Name  string
	Type  string
	Value string
	TTL   int
}

type Zone struct {
	records map[string]*Record
}

// TODO: define Evict(name string) to remove the entry for name from the cache
func (c *Cache) Evict(name string) {
	_, ok := c.entries[name]
	if ok {
		delete(c.entries, name)
	}
}

// TODO: define Count() int to return the number of entries in the cache
func (c *Cache) Count() int {
	count := 0
	for range c.entries {
		count++
	}
	return count
}

func NewZone() *Zone {
	return &Zone{
		records: make(map[string]*Record),
	}
}

func (z *Zone) AddRecord(r *Record) {
	z.records[r.Name] = r
}

func (z *Zone) GetRecord(name string) (*Record, bool) {
	r, ok := z.records[name]
	return r, ok
}

func (z *Zone) RemoveRecord(name string) bool {
	_, ok := z.records[name]
	if ok {
		delete(z.records, name)
	}
	return ok
}

func (z *Zone) ListRecords() []*Record {
	records := make([]*Record, 0, len(z.records))
	for _, r := range z.records {
		records = append(records, r)
	}
	return records
}

type CacheEntry struct {
	Record  *Record
	AddedAt int
}

type Cache struct {
	entries map[string]*CacheEntry
	now     func() int
}

func NewCache(now func() int) *Cache {
	return &Cache{
		entries: make(map[string]*CacheEntry),
		now:     now,
	}
}

func (c *Cache) Store(record *Record) {
	c.entries[record.Name] = &CacheEntry{
		Record:  record,
		AddedAt: c.now(),
	}
}

func (c *Cache) Lookup(name string) (*Record, bool) {
	entry, ok := c.entries[name]
	if !ok {
		return nil, false
	}
	if c.now() >= entry.AddedAt+entry.Record.TTL {
		return nil, false
	}
	return entry.Record, true
}

func main() {
	zone := NewZone()
	api := &Record{Name: "api.example.com", Type: "A", Value: "93.184.216.34", TTL: 300}
	zone.AddRecord(api)

	clock := 1000
	cache := NewCache(func() int { return clock }) // fake clock for the demo
	cache.Store(api)

	fmt.Printf("Count: %d\n", cache.Count())
	cache.Evict("api.example.com")
	fmt.Printf("After evict: %d\n", cache.Count())

	// Pointer sharing demo
	cache.Store(api)
	api.Value = "93.184.216.35"
	clock = 1100
	if r, ok := cache.Lookup("api.example.com"); ok {
		fmt.Printf("Cache sees update: %s\n", r.Value)
	}
}
