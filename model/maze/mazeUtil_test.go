package maze

import (
	"math/rand"
	"testing"
)

func TestCoordToDataPosBasic(t *testing.T) {
	exp := 15

	x, y, colDims := 3, 2, 6

	res := CoordToDataPos(x, y, colDims)

	if res != exp {
		t.Error("expected:", exp, "got:", res)
	}
}

func TestCoordToDataPosLower(t *testing.T) {
	x, y, colDims := 0, 0, 3

	exp := 0

	res := CoordToDataPos(x, y, colDims)

	if res != exp {
		t.Error("expected:", exp, "got:", res)
	}
}

func TestCoordToDataPosUpper(t *testing.T) {
	x, y, colDims := 0, 0, 3

	exp := 0

	res := CoordToDataPos(x, y, colDims)

	if res != exp {
		t.Error("expected:", exp, "got:", res)
	}
}

func TestTargetZoneBasic(t *testing.T) {
	for i := 0; i < 100; i++ {
		r, c := rand.Intn(10)+10, rand.Intn(100)+10
		println(i, "Row", r, "col", c)
		target := GenTargetZone(r, c)

		if target.x < int(float64(c)*0.15) || target.x > int(float64(c)*0.85) {
			t.Error(i, "x", target.x, "outside bounds", int(float64(c)*0.15), int(float64(c)*0.85))
		}
		if target.y < int(float64(r)*0.15) || target.y > int(float64(r)*0.85) {
			t.Error(i, "y", target.y, "outside bounds", int(float64(r)*0.15), int(float64(r)*0.85))
		}
	}
}
