package main

import "fmt"

// TODO: Define the Record struct with Name (string), Type (string), Value (string), TTL (int)
type Record struct {
	Name  string
	Type  string
	Value string
	TTL   int
}

// TODO: Define the Zone struct with records (map[string]*Record)
type Zone struct {
	records map[string]*Record
}

// TODO: NewZone() *Zone - create and return a new zone with initialized map
func NewZone() *Zone {
	return &Zone{records: make(map[string]*Record)}
}

// TODO: (z *Zone) AddRecord(r *Record) - store record pointer using r.Name as key
func (z *Zone) AddRecord(r *Record) {
	z.records[r.Name] = r
}

func main() {
	zone := NewZone()
	zone.AddRecord(&Record{Name: "api.example.com", Type: "A", Value: "93.184.216.34", TTL: 300})
	zone.AddRecord(&Record{Name: "cdn.example.com", Type: "CNAME", Value: "d1234.cloudfront.net", TTL: 3600})
	fmt.Printf("Zone has %d records\n", len(zone.records))
}
