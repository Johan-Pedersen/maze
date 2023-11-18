package maze

/*
type fil

*/

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

type Coordinate struct {
	// Hvorfor skal alle disse vœre public
	X, Y int
}

type Maze struct {
	Maze *mat.Dense `json:"maze_track,omitempty"`

	Target Coordinate `json:"target,omitempty"`

	Paths []Coordinate `json:"paths,omitempty"`

	YBound, XBound int
}

/*
Defines the 2d vector (x y),
as the vector going x steps on the x-axis and y steps on the y-axis
*/
type Vector struct {
	X, Y int
}

/*
Create vector
*/
func NewVector(head, target Coordinate) Vector {
	return Vector{target.X - head.X, target.Y - head.Y}
}

/*
Create normalized vector, with ||vec||_2 =1
*/
func NewNormVector(head, target Coordinate) Vector {
	vec := NewVector(head, target)
	return Vector{(vec.X) / int(Norm(vec)), vec.Y / int(Norm(vec))}
}

type StepDirection int

const (
	Left  StepDirection = 0
	Right StepDirection = 1
	Up    StepDirection = 2
	Down  StepDirection = 3
)

func (path Coordinate) String() string {
	return fmt.Sprintf("X: %d, Y: %d", path.X, path.Y)
}
