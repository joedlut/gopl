package main

import (
	"fmt"
	"image/color"
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

type ColoredPoint struct {
	Point
	Color color.Color
}

func main() {
	var cp ColoredPoint
	cp.x = 1
	fmt.Println(cp.Point.x)
	cp.y = 90
	fmt.Println(cp.y)
	cp.ScaleBy(2)
	fmt.Println(cp.x, cp.y)

	var cp1 = ColoredPoint{Point{1, 2}, color.RGBA{0, 1, 1, 1}}
	fmt.Println(cp1)
	// cp1.Distance(cp)  错误
	fmt.Println(cp1.Distance(cp.Point))
}
