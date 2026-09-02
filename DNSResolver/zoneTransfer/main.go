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

// TODO: Define Transfer struct with Records ([]Record), Source (string), Status (string)
type Transfer struct {
	Records []Record
	Source string
	Status string
}

// TODO: (dst *Zone) TransferFrom(src *Zone, names []string) (*Transfer, error)
//       1. For each name: look up in src, validate, add to dst, track for rollback
//       2. On failure: remove all added records from dst, return error
//       3. On success: create Transfer with value copies, Source "primary", Status "complete"

func (dst *Zone) TransferFrom(src *Zone, names []string) (*Transfer, error) {
	recordList := []Record{}
	for i, name := range names {
		r, ok := src.GetRecord(name)
		if !ok {
			for j := 0 ; j < i ; j++ {
				dst.RemoveRecord(names[j])
			}
			return nil, fmt.Errorf("record not found: %s",name)
		}
		err := validateRecord(r)
		if err != nil{
			for j := 0 ; j < i ; j++ {
				dst.RemoveRecord(names[j])
			}
			return nil, fmt.Errorf("invalid record %s: %w",name, err)
		}
		dst.AddRecord(r)
		recordList = append(recordList,*r)
	}
	return &Transfer{
		Records : recordList,
		Source : "primary",
		Status : "complete",
	},nil
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

func detectCycle(zone *Zone, name string) bool {
	visited := map[string]bool{}

	current := name
	for {
		if visited[current] {
			return true
		}
		visited[current] = true

		record, ok := zone.GetRecord(current)
		if !ok {
			return false
		}
		if record.Type != "CNAME" {
			return false
		}
		current = record.Value
	}
}

func (z *Zone) ValidateAll() []error {
	var errs []error
	for _, r := range z.records {
		if err := validateRecord(r); err != nil {
			errs = append(errs, err)
		}
	}
	for _, r := range z.records {
		if r.Type == "CNAME" && detectCycle(z, r.Name) {
			errs = append(errs, fmt.Errorf("CNAME cycle detected: %s", r.Name))
		}
	}
	return errs
}

func main() {
	primary := NewZone()
	primary.AddRecord(&Record{Name: "api.example.com", Type: "A", Value: "93.184.216.34", TTL: 300})
	primary.AddRecord(&Record{Name: "cdn.example.com", Type: "CNAME", Value: "d1234.cloudfront.net", TTL: 3600})

	backup := NewZone()
	transfer, err := backup.TransferFrom(primary, []string{"api.example.com", "cdn.example.com"})
	if err != nil {
		fmt.Println("Transfer failed:", err)
		return
	}
	fmt.Printf("Transfer %s: %d records from %s\n", transfer.Status, len(transfer.Records), transfer.Source)
}