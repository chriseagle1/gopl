package main

import "fmt"

type Point struct {
	x, y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	w1 := Wheel{Circle{Point{8, 8}, 5}, 20}

	w2 := Wheel{
		Circle:Circle{
			Point:Point{
				x: 8,
				y: 8,
			},
			Radius: 5,
		},
		Spokes: 20,
	}

	fmt.Printf("%#v\n", w1)
	fmt.Printf("%#v\n", w2)

	w1.x = 42
	w2.Radius = 10

	fmt.Printf("%#v\n", w1)
	fmt.Printf("%#v\n", w2)
}
