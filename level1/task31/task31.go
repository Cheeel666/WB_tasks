package task31

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func construct(x float64, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func distance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}

func Run() {
	a := construct(1, 1)
	b := construct(3, 1)
	fmt.Println("A:", a, "\nB:", b)
	fmt.Println("Distance between dots:", distance(a, b))
}
