package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func Distance(p, q Point) float64 {
	return math.Hypot(p.x-q.x, p.y-q.y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.x-q.x, p.y-q.y)
}

func (p *Point) ScaleBy(factor float64) {
	p.x *= factor
	p.y *= factor
}

type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}
func main() {
	var p Point = Point{0, 0}
	var q Point = Point{3, 4}

	var q1 *Point = &Point{3, 4}
	var x Point = Point{6, 8}
	fmt.Println(p.Distance(q))
	fmt.Println(Distance(p, q))
	path := Path{p, q, x}
	fmt.Println(path.Distance())
	factor := 0.5
	q.ScaleBy(factor)
	q1.ScaleBy(factor)
	fmt.Println(q)
	fmt.Println(*q1)
}
