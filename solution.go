package square

import (
	"math"
)

type Sides int

const (
	SidesCircle Sides = iota
	_
	_
	SidesTriangle
	SidesSquare
)

func CalcSquare(sideLen float64, sidesNum Sides) float64 {
	var sideLenSquare = math.Pow(sideLen, 2)
	switch sidesNum {
	case 3:
		return sideLenSquare * math.Sqrt(3) / 4
	case 4:
		return sideLenSquare
	case 0:
		return math.Pi * sideLenSquare
	}
	return 0
}