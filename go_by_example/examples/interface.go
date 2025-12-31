package examples

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (s square) perimeter() float64 {
	return 4 * s.side
}

func DemoInterface() {
	var g geometry

	g = circle{radius: 5}
	fmt.Printf("Circle - Area: %.2f, Perimeter: %.2f\n", g.area(), g.perimeter())

	g = square{side: 4}
	fmt.Printf("Square - Area: %.2f, Perimeter: %.2f\n", g.area(), g.perimeter())
}