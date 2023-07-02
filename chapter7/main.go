package main

import (
	"fmt"
	"math"
)

func main() {
	// var circle Circle
	// c := *Circle{0, 0, 5}
	c := Circle{0, 0, 5}
	fmt.Println("Circle area=", c.area())
	r := Rectangle{1, 4, 2, 2}
	fmt.Println("Rectangle area=", r.area())
	totalShapesArea := totalArea(&c, &r)
	fmt.Println("Total shapes area=", totalShapesArea)
	totalShapesPerimeter := totalPerimeter(&c, &r)
	fmt.Println("Total shapes perimeter=", totalShapesPerimeter)

	e := Employee{}
	e.name = "Amr"
	e.talk()
}

type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

type Rectangle struct {
	x1, x2, y1, y2 float64
}

func (r *Rectangle) area() float64 {
	length, width := distance(r.x1, r.y1, r.x2, r.y2)
	return length * width
}

func (r *Rectangle) perimeter() float64 {
	length, width := distance(r.x1, r.y1, r.x2, r.y2)
	return 2 * (length + width)
}

func distance(x1, y1, x2, y2 float64) (l, w float64) {
	l = math.Abs(x2 - x1)
	w = math.Abs(y2 - y1)
	return l, w
}

func totalArea(shapes ...Shape) float64 {
	total := 0.0
	for _, value := range shapes {
		total += value.area()
	}

	return total
}

func totalPerimeter(shapes ...Shape) float64 {
	total := 0.0
	for _, value := range shapes {
		total += value.perimeter()
	}

	return total
}

type Person struct {
	name string
	age  int
}

func (p *Person) talk() {
	fmt.Println("Hi, my name is", p.name)
}

type Employee struct {
	//==> (is a person) relationship, this filed is called anonymous filed
	Person
	//==> (has an id) relationship
	id int64
}
