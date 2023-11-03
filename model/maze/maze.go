package maze

/*
type fil

*/

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

// Hvorfor skal alle disse vœre public
type TargetCoordinate struct {
	X, Y int `json:"x,omitempty"`
}

type PathCoordinate struct {
	X, Y int
}

type Maze struct {
	Maze *mat.Dense `json:"maze_track,omitempty"`

	Target TargetCoordinate `json:"target,omitempty"`

	Paths []PathCoordinate `json:"paths,omitempty"`

	YBound, XBound int
}

type TargetVector struct {
	// BaseX og BaseY skal nok vœre optional. Fordi de afhœnger jo af hvilket head man har

	// Vi skal også lige overveje hvad der skal ske når vi opdatere en pathcoordinate, fordi den skal jo følge med. Så det skal vœre et objekt for hver
	BaseX, BaseY, HeadX, HeadY int
}

type StepDirection int

const (
	Left  StepDirection = 0
	Right StepDirection = 1
	Up    StepDirection = 2
	Down  StepDirection = 3
)

func (path PathCoordinate) String() string {
	return fmt.Sprintf("X: %d, Y: %d", path.X, path.Y)
}


