package maze

import "testing"

func TestIsDirValidXLow(t *testing.T) {
	exp := false

	head := coordinate{0, 2}

	xBound, yBound := 2, 3

	left := stepDirection(0)

	res := isDirValid(left, head, xBound, yBound)

	if res != exp {
		t.Error("expected: ", exp, "got:", res)
	}
}
