package task28

import (
	"fmt"
	"math"
)

// круглое отверстие
type RoundHole struct {
	radius float64
}

func (rh *RoundHole) getRadius() float64 {
	return rh.radius
}

func (rh *RoundHole) fits(radius float64) bool {
	return radius <= rh.radius
}

// Круглые колышки
type RoundPeg struct {
	radius float64
}

func (rp *RoundPeg) getRadius() float64 {
	return rp.radius
}

// Квадратные колышки
type SquarePeg struct {
	width float64
}

func (sp *SquarePeg) getWidth() float64 {
	return sp.width
}

// Адаптер для квадратный колышков
type SquarePegAdapter struct {
	peg SquarePeg
}

func (spa *SquarePegAdapter) getRadius() float64 {
	return float64(spa.peg.width) * math.Sqrt(2) / 2
}
func Run() {
	var hole = RoundHole{radius: 10}
	var sqPeg = SquarePeg{width: 13}
	var peg = SquarePegAdapter{sqPeg}
	fmt.Println(hole)
	fmt.Println(peg.getRadius())
	fmt.Println("Fits:", hole.fits(peg.getRadius()))
}
