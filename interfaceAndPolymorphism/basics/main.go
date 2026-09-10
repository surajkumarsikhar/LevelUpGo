package main

import "math"

// TODO: Define the Shape interface with Area() and Perimeter() methods
type Shape interface {
	Area() float64
	Perimeter() float64
}

// TODO: Define Circle struct with Radius float64
type Circle struct {
	Radius float64
}

// TODO: Implement Area() and Perimeter() methods for Circle
func (c Circle) Area() float64 {
	return (math.Pi * c.Radius * c.Radius)
}
func (c Circle) Perimeter() float64 {
	return (2 * math.Pi * c.Radius)
}

// TODO: Define Rectangle struct with Width, Height float64
type Rectangle struct {
	Width  float64
	Height float64
}

// TODO: Implement Area() and Perimeter() methods for Rectangle
func (r Rectangle) Area() float64 {
	return (r.Width * r.Height)
}
func (r Rectangle) Perimeter() float64 {
	return (2 * (r.Width + r.Height))
}

// TODO: Implement TotalArea(shapes []Shape) float64
func TotalArea(shapes []Shape) float64 {
	totalArea := 0.0
	for _, shape := range shapes {
		totalArea += shape.Area()
	}
	return totalArea
}

// TODO: Implement TotalPerimeter(shapes []Shape) float64
func TotalPerimeter(shapes []Shape) float64 {
	totalPerimeter := 0.0
	for _, shape := range shapes {
		totalPerimeter += shape.Perimeter()
	}
	return totalPerimeter
}

func main() {}
