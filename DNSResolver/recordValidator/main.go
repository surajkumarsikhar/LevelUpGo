package main

import "fmt"

type Record struct {
	Name  string
	Type  string
	Value string
	TTL   int
}

// TODO: validateRecord(r *Record) error
//       Check in order: nil, empty Name, invalid Type, negative TTL, empty Value
func validateRecord(r *Record) error {
	if r == nil {
		return fmt.Errorf("nil record")
	}
	if r.Name == "" {
		return fmt.Errorf("empty name")
	}
	if r.Type != "A" && r.Type != "CNAME" && r.Type != "MX" && r.Type != "TXT" {
		return fmt.Errorf("unknown type: %s", r.Type)
	}
	if r.TTL < 0 {
		return fmt.Errorf("negative TTL")
	}
	if r.Value == "" {
		return fmt.Errorf("empty value")
	}
	return nil
}

type Zone struct {
	records map[string]*Record
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

func (c *Cache) Evict(name string) {
	delete(c.entries, name)
}

func (c *Cache) Count() int {
	return len(c.entries)
}

func main() {
	err := validateRecord(nil)
	fmt.Println("nil record:", err)

	err = validateRecord(&Record{Name: "api.example.com", Type: "A", Value: "93.184.216.34", TTL: 300})
	fmt.Println("valid record:", err)

	err = validateRecord(&Record{Name: "x.com", Type: "AAAA", Value: "::1", TTL: 300})
	fmt.Println("bad type:", err)
}
