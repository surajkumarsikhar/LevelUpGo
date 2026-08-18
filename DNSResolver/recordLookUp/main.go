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

// TODO: (z *Zone) GetRecord(name string) (*Record, bool) - look up record by name
func (z *Zone) GetRecord(name string) (*Record, bool) {
	r, ok := z.records[name]
	return r, ok
}

// TODO: (z *Zone) ListRecords() []*Record - return all records
func (z *Zone) ListRecords() []*Record {
	r := []*Record{}
	for _, record := range z.records {
		r = append(r, record)
	}
	return r
}

// TODO: (z *Zone) RemoveRecord(name string) bool - remove record, return true if found
func (z *Zone) RemoveRecord(name string) bool {
	_, ok := z.records[name]
	delete(z.records, name)
	return ok
}

func NewZone() *Zone {
	return &Zone{
		records: make(map[string]*Record),
	}
}

func (z *Zone) AddRecord(r *Record) {
	z.records[r.Name] = r
}

func main() {
	zone := NewZone()
	zone.AddRecord(&Record{Name: "api.example.com", Type: "A", Value: "93.184.216.34", TTL: 300})
	zone.AddRecord(&Record{Name: "cdn.example.com", Type: "CNAME", Value: "d1234.cloudfront.net", TTL: 3600})

	if r, ok := zone.GetRecord("api.example.com"); ok {
		fmt.Printf("%s -> %s\n", r.Name, r.Value)
	}

	fmt.Printf("Records: %d\n", len(zone.ListRecords()))
	zone.RemoveRecord("cdn.example.com")
	fmt.Printf("After remove: %d\n", len(zone.ListRecords()))
}
