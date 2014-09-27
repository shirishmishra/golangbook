package main

import ("fmt"; "math")

type Shape interface {
	perimeter() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	length, breadth float64
}
func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r *Rectangle) perimeter() float64 {
	return 2 * (r.length + r.breadth)
}
func main() {
		c := Circle{5}
		r := Rectangle{4, 3}

		show(&c, &r)
}

func show(sh ...Shape) {
		for _, s := range sh {
			fmt.Println(s.perimeter())
		}
}
