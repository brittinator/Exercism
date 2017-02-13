//package triangle: this took 35 minutes
package triangle

import (
	"math"
)

const (
	testVersion      = 3
	NaT         Kind = "Not a triangle"
	Equ         Kind = "Equalateral triangle"
	Iso         Kind = "Isosceles triangle"
	Sca         Kind = "Scalene triangle"
)

//Kind is the type of triangle returned
//from KindFromSides function
type Kind string

//KindFromSides returns the type of triangle given the
//sides, or NaT if it isn't a triangle
func KindFromSides(a, b, c float64) Kind {
	if isTriangle([]float64{a, b, c}) {
		if a == b && b == c {
			return Equ
		}
		if a == b || b == c || a == c {
			return Iso
		}
		return Sca
	}
	return NaT
}

//isTriangle returns true if given sides can be a triangle
func isTriangle(sides []float64) bool {
	for _, side := range sides {
		if math.IsNaN(side) {
			return false
		}
		if math.IsInf(side, 0) {
			return false
		}
	}

	if sides[0]+sides[1] >= sides[2] && sides[1]+sides[2] >= sides[0] && sides[0]+sides[2] >= sides[1] {
		for _, side := range sides {
			if side <= 0 {
				return false
			}
		}
	} else {
		return false
	}
	return true
}
