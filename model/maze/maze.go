package maze

import (
	"fmt"
	"math/rand"

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

type TargetVektor struct {
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

var yes = 0

func (path PathCoordinate) String() string {
	return fmt.Sprintf("X: %d, Y: %d", path.X, path.Y)
}

func NewMaze(r, c int) Maze {
	data := make([]float64, r*c)
	for i := 0; i < r*c; i++ {
		data[i] = 1
	}

	// Generate 2x2 target zone
	targetX, targetY := TargetZone(r, c)

	data[CoordToDataPos(targetX, targetY, c)] = 2
	data[CoordToDataPos(targetX+1, targetY, c)] = 2
	data[CoordToDataPos(targetX, targetY+1, c)] = 2
	data[CoordToDataPos(targetX+1, targetY+1, c)] = 2

	mz := &Maze{
		Maze:   mat.NewDense(r, c, data),
		Target: TargetCoordinate{targetX, targetY},
		Paths:  []PathCoordinate{},
		YBound: r,
		XBound: c,
	}

	ripple(mz.Maze, mz.Target)

	createPath(mz)

	return *mz
}

func ripple(maze *mat.Dense, target TargetCoordinate) {

}

func validateDir(dir StepDirection, head PathCoordinate, xAxisBound, yAxisBound int) bool {
	return (dir == Left && head.X == 0) ||
		(dir == Right && head.X == xAxisBound-1) ||
		(dir == Up && head.Y == yAxisBound-1) ||
		(dir == Down && head.Y == 0)
}

func createPath(mz *Maze) {
	// Init first position
	mz.Paths = append(mz.Paths, PathCoordinate{0, 0})

	mzTrack := *mz.Maze
	mzTrack.Set(0, 0, float64(0))

	// direction := [4] StepDirection {Left, Right, Up, Down}

	println("left:", Left)
	println("Right", Right)
	println("Up:", Up)
	println("Down:", Down)
	fmt.Println("dims", mz.XBound, mz.YBound)
	// Tager 10 skridt
	for i := 0; i < 20; i++ {

		println("******************", i)

		dir := StepDirection(rand.Intn(4))

		// Denne kunne man godt lave til en pointer. Så man kunne skrive direkte til memory location
		head := mz.Paths[0]

		println("dir: ", dir)

		if validateDir(dir, head, mz.XBound, mz.YBound) {
			continue
		}

		switch dir {
		case Left:
			if mzTrack.At(head.Y, head.X-1) != 0 {
				head.X = head.X - 1
			}
		case Right:
			if mzTrack.At(head.Y, head.X+1) != 0 {
				head.X = head.X + 1
			}
		case Up:
			if mzTrack.At(head.Y+1, head.X) != 0 {
				head.Y = head.Y + 1
			}
		case Down:
			if mzTrack.At(head.Y-1, head.Y) != 0 {
				head.Y = head.Y - 1
			}
		}

		mz.Maze.Set(head.Y, head.X, 0)

		fmt.Println("head", head)

		mz.Paths[0] = head

		PrintMaze(*mz)

	}
}
