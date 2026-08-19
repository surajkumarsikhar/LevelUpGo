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

// TODO: detectCycle(zone *Zone, name string) bool
//       Follow CNAME chain with a visited map. Return true if cycle found.
func detectCycle(zone *Zone, name string) bool {
	seen := make(map[string]bool)
	for !seen[name] {
		seen[name] = true
		rec, ok := zone.records[name]
		if !ok {
			return false
		}
		if rec.Type != "CNAME" {
			return false
		}
		name = zone.records[name].Value

	}
	return true
}

// TODO: (z *Zone) ValidateAll() []error
//       Run validateRecord on all records, then detectCycle on all CNAMEs.
//       Collect all errors into a single slice.
func (z *Zone) ValidateAll() []error {
	errorSlice := []error{}
	for _, r := range z.records {
		err := validateRecord(r)
		if err != nil {
			errorSlice = append(errorSlice, err)
		}
	}

	for _, r := range z.records {
		cycleDetected := detectCycle(z, r.Name)
		if cycleDetected {
			errorSlice = append(errorSlice, fmt.Errorf("CNAME cycle detected: %s", r.Name))
		}
	}
	return errorSlice
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

func main() {
	zone := NewZone()
	zone.AddRecord(&Record{Name: "api.example.com", Type: "A", Value: "93.184.216.34", TTL: 300})
	zone.AddRecord(&Record{Name: "www.example.com", Type: "CNAME", Value: "cdn.example.com", TTL: 300})
	zone.AddRecord(&Record{Name: "cdn.example.com", Type: "CNAME", Value: "www.example.com", TTL: 300})

	fmt.Printf("Cycle at www? %v\n", detectCycle(zone, "www.example.com"))
	fmt.Printf("Cycle at api? %v\n", detectCycle(zone, "api.example.com"))

	errs := zone.ValidateAll()
	for _, err := range errs {
		fmt.Println("Error:", err)
	}
}
