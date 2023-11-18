package maze

/*
type fil

*/

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

type coordinate struct {
	// Hvorfor skal alle disse vœre public
	X, Y int
}

type Maze struct {
	Maze *mat.Dense `json:"maze_track,omitempty"`

	Target coordinate `json:"target,omitempty"`

	Paths []coordinate `json:"paths,omitempty"`

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
func NewVector(head, target coordinate) Vector {
	return Vector{target.X - head.X, target.Y - head.Y}
}

/*
Create normalized vector, with ||vec||_2 =1
*/
func NewNormVector(head, target coordinate) Vector {
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

func (path coordinate) String() string {
	return fmt.Sprintf("X: %d, Y: %d", path.X, path.Y)
}
