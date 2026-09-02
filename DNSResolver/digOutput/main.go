package main

import (
	"fmt"
	"strings"
)

type Record struct {
	Name  string
	Type  string
	Value string
	TTL   int
}

type Zone struct {
	records map[string]*Record
}

// TODO: formatRecord(r *Record) string
//
//	Return one record as a single line: name (with trailing dot),
//	TTL, the class IN, type, and value, padded into aligned columns.
func formatRecord(r *Record) string {
	return fmt.Sprintf("%-24s %-6d IN  %-6s %s\n", r.Name+".", r.TTL, r.Type, r.Value)
}

// TODO: resolve(zone *Zone, names []string) (string, error)
//       1. Validate zone with ValidateAll(), return first error if any
//       2. Build output starting with ";; ANSWER SECTION:\n"
//       3. Look up each name, return NXDOMAIN error if not found
//       4. Append formatted records and summary line

func resolve(zone *Zone, names []string) (string, error) {
	err := zone.ValidateAll()
	for _, e := range err {
		if e != nil {
			return "", e
		}
	}
	output := []string{}

	for _, name := range names {
		r, ok := zone.records[name]
		if !ok {
			return "", fmt.Errorf("NXDOMAIN: %s", name)
		}
		output = append(output, formatRecord(r))
	}

	outputStr := fmt.Sprintf(";; ANSWER SECTION:\n%s\n;; RECORDS: %d\n;; STATUS: NOERROR\n", strings.Join(output, ""), len(names))
	return outputStr, nil
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

type Transfer struct {
	Records []Record
	Source  string
	Status  string
}

func (dst *Zone) TransferFrom(src *Zone, names []string) (*Transfer, error) {
	var added []string

	for _, name := range names {
		record, ok := src.GetRecord(name)
		if !ok {
			for _, a := range added {
				dst.RemoveRecord(a)
			}
			return nil, fmt.Errorf("record not found: %s", name)
		}

		if err := validateRecord(record); err != nil {
			for _, a := range added {
				dst.RemoveRecord(a)
			}
			return nil, fmt.Errorf("invalid record %s: %w", name, err)
		}

		dst.AddRecord(record)
		added = append(added, name)
	}

	records := make([]Record, 0, len(added))
	for _, name := range added {
		r, _ := dst.GetRecord(name)
		records = append(records, *r)
	}

	return &Transfer{
		Records: records,
		Source:  "primary",
		Status:  "complete",
	}, nil
}

func main() {
	zone := NewZone()
	zone.AddRecord(&Record{Name: "api.example.com", Type: "A", Value: "93.184.216.34", TTL: 300})
	zone.AddRecord(&Record{Name: "cdn.example.com", Type: "CNAME", Value: "d1234.cloudfront.net", TTL: 3600})
	zone.AddRecord(&Record{Name: "mail.example.com", Type: "MX", Value: "10 smtp.example.com", TTL: 7200})

	output, err := resolve(zone, []string{"api.example.com", "cdn.example.com", "mail.example.com"})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Print(output)
}
