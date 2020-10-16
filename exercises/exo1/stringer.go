package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	fmt.Stringer
}

type Circle struct {
	r float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle : length = %f, width = %f", r.Width, r.Length)
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle of radius %f", c.r)
}

func main() {
	r := &Rectangle{2, 4}
	c := &Circle{4}

	shapes := []Shape{r, c}

	for _, s := range shapes {
		fmt.Println(s)
		fmt.Println("area =", s.Area())
	}
}
