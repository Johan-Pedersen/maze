package maze_test

import (
	"math/rand"
	"testing"

	"maze/model/maze"
)

func TestCoordToDataPosBasic(t *testing.T) {
	exp := 15

	x, y, colDims := 3, 2, 6

	res := maze.CoordToDataPos(x, y, colDims)

	if res != exp {
		t.Error("expected:", exp, "got:", res)
	}
}

func TestCoordToDataPosLower(t *testing.T) {
	x, y, colDims := 0, 0, 3

	exp := 0

	res := maze.CoordToDataPos(x, y, colDims)

	if res != exp {
		t.Error("expected:", exp, "got:", res)
	}
}

func TestCoordToDataPosUpper(t *testing.T) {
	x, y, colDims := 0, 0, 3

	exp := 0

	res := maze.CoordToDataPos(x, y, colDims)

	if res != exp {
		t.Error("expected:", exp, "got:", res)
	}
}

func TestTargetZoneBasic(t *testing.T) {
	for i := 0; i < 100; i++ {
		r, c := rand.Intn(10)+10, rand.Intn(100)+10
		println(i, "Row", r, "col", c)
		x, y := maze.TargetZone(r, c)

		if x < int(float64(c)*0.15) || x > int(float64(c)*0.85) {
			t.Error(i, "x", x, "outside bounds", int(float64(c)*0.15), int(float64(c)*0.85))
		}
		if y < int(float64(r)*0.15) || y > int(float64(r)*0.85) {
			t.Error(i, "y", y, "outside bounds", int(float64(r)*0.15), int(float64(r)*0.85))
		}
	}
}
