package main

import "fmt"

// 匿名嵌套
type Point struct {
	x, y int
}

type Circle struct {
	radius int
	Point
}

type Wheel struct {
	Circle
	spokes int
}

func main() {
	//var w Wheel
	//w.x = 10
	//w.y = 30
	w1 := Wheel{Circle{5, Point{2, 3}}, 10}
	fmt.Printf("%#v\n", w1)

	w2 := Wheel{
		Circle: Circle{
			Point:  Point{x: 8, y: 8},
			radius: 45,
		},
		spokes: 12,
	}

	fmt.Printf("%#v\n", w2)
}
