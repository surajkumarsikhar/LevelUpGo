package main

import "fmt"

// Define your Address struct here
type Address struct {
	Street  string
	City    string
	Country string
}

// Define your Employee struct here (embed Address)
type Employee struct {
	Address
	Name string
	Role string
}

// Write your NewEmployee function here
func NewEmployee(name, role, street, city, country string) Employee {
	return Employee{
		Address: Address{
			Street:  street,
			City:    city,
			Country: country,
		},
		Name: name,
		Role: role,
	}
}

// Write your FullAddress method here
func (e Employee) FullAddress() string {
	return fmt.Sprintf("%s, %s, %s", e.Street, e.City, e.Country)
}

func main() {
	e := NewEmployee("Alice", "Engineer", "123 Main St", "Portland", "USA")
	fmt.Println(e.City)
	fmt.Println(e.FullAddress())
}
