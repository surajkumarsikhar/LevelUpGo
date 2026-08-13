package main

import "fmt"

type Sensor struct {
	Name    string
	Reading float64
}

// Write your Status method here (value receiver)
func (s Sensor) Status() (msg string) {
	msg = fmt.Sprintf("%s: %.1f°C", s.Name, s.Reading)
	return
}

// Write your Record method here (pointer receiver)
func (s *Sensor) Record(value float64) {
	s.Reading = value
}

func main() {
	s := Sensor{Name: "engine", Reading: 0}
	s.Record(92.5)
	fmt.Println(s.Status())
	s.Record(100.3)
	fmt.Println(s.Status())
}
