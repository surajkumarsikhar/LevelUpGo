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

// TODO: Define CacheEntry struct with Record (*Record) and AddedAt (int)
type CacheEntry struct {
	Record  *Record
	AddedAt int
}

// TODO: Define Cache struct with entries (map[string]*CacheEntry) and now (func() int)
type Cache struct {
	entries map[string]*CacheEntry
	now     func() int
}

// TODO: NewCache(now func() int) *Cache - initialized map, now set from the parameter
func NewCache(now func() int) *Cache {
	return &Cache{
		entries: make(map[string]*CacheEntry),
		now:     now,
	}
}

// TODO: (c *Cache) Store(record *Record) - create entry with AddedAt = c.now()
func (c *Cache) Store(record *Record) {
	ce := &CacheEntry{
		Record:  record,
		AddedAt: c.now(),
	}
	c.entries[record.Name] = ce
}

// TODO: (c *Cache) Lookup(name string) (*Record, bool) - return record if fresh
func (c *Cache) Lookup(name string) (*Record, bool) {
	ce, ok := c.entries[name]
	if !ok {
		return nil, false
	} else if exp := ce.AddedAt + ce.Record.TTL; exp <= c.now() {
		return nil, false
	}
	return ce.Record, true
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

func main() {
	api := &Record{Name: "api.example.com", Type: "A", Value: "93.184.216.34", TTL: 300}

	clock := 1000
	cache := NewCache(func() int { return clock }) // inject a clock we control
	cache.Store(api)

	clock = 1100 // 100s later
	if r, ok := cache.Lookup("api.example.com"); ok {
		fmt.Printf("Cache hit: %s -> %s\n", r.Name, r.Value)
	}

	clock = 1300 // TTL passed
	if _, ok := cache.Lookup("api.example.com"); !ok {
		fmt.Println("Cache miss: record expired")
	}
}
