package maze

/*
type fil

*/

import (
	"fmt"
	"strconv"

	"gonum.org/v1/gonum/mat"
)

type coordinate struct {
	x, y int
}

type Maze struct {
	Maze *mat.Dense

	Target coordinate

	paths []coordinate

	yBound, xBound int
}

/*
Defines the 2d vector (x y),
as the vector going x steps on the x-axis and y steps on the y-axis
*/
type vector struct {
	x, y float64
}

func (vec vector) String() string {
	return "(" +
		strconv.FormatFloat(vec.x, 'f', -1, 64) + " " +
		strconv.FormatFloat(vec.y, 'f', -1, 64) +
		")"
}

/*
Create vector
*/
func newVector(head, target coordinate) vector {
	return vector{float64(target.x - head.x), float64(target.y - head.y)}
}

/*
Create normalized vector, with ||vec||_2 =1
*/
func newNormVector(head, target coordinate) vector {
	vec := newVector(head, target)
	return vector{(vec.x) / norm(vec), vec.y / norm(vec)}
}

type stepDirection int

const (
	Left  stepDirection = 0
	Right stepDirection = 1
	Up    stepDirection = 2
	Down  stepDirection = 3
)

func (dir stepDirection) String() string {
	switch dir {
	case 0:
		return "Left"
	case 1:
		return "Right"
	case 2:
		return "Up"
	default:
		return "Down"
	}
}

func (coord coordinate) String() string {
	return fmt.Sprintf("X: %d, Y: %d", coord.x, coord.y)
}

func (coord coordinate) equals(other coordinate) bool {
	return other.x == coord.x && other.y == coord.y
}
