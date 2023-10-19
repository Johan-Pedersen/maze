package mazeGen

import (
	"fmt"
	"math/rand"
	"strings"

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
	MazeTrack *mat.Dense `json:"maze_track,omitempty"`

	Target TargetCoordinate `json:"target,omitempty"`

	Paths []PathCoordinate `json:"paths,omitempty"`
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

func GenerateMaze(r, c int) Maze {
	data := make([]float64, r*c)
	for i := 0; i < r*c; i++ {
		data[i] = 1
	}

	// Set target location
	targetYLowerBound := 0.25 * float64(r)
	targetXLowerBound := 0.25 * float64(c)

	targetY := rand.Intn(r-1) + int(targetYLowerBound)
	targetX := rand.Intn(c-1) + int(targetXLowerBound)

	data[(targetY-1)*c+targetX] = 2
	data[(targetY)*c+targetX] = 2
	data[(targetY-1)*c+targetX+1] = 2
	data[(targetY)*c+targetX+1] = 2

	mz := Maze{
		MazeTrack: mat.NewDense(r, c, data),
		Target:    TargetCoordinate{targetX, targetY},
		Paths:     []PathCoordinate{},
	}

	return createPath(mz)
}

func createPath(mz Maze) Maze {
	// Init first position
	mz.Paths = append(mz.Paths, PathCoordinate{0, 0})

	mzTrack := *mz.MazeTrack
	mzTrack.Set(0, 0, float64(0))

	// direction := [4] StepDirection {Left, Right, Up, Down}

	println("left:", Left)
	println("Right", Right)
	println("Up:", Up)
	println("Down:", Down)
	yAxisBound, xAxisBound := mz.MazeTrack.Dims()
	println("yaxis bound:", yAxisBound, "xaxixbound", xAxisBound)
	r, c := mz.MazeTrack.Dims()
	fmt.Println("dims", r, c)
	// Tager 10 skridt
	for i := 0; i < 20; i++ {

		println("******************", i)

		dir := StepDirection(rand.Intn(4))

		// Denne kunne man godt lave til en pointer. Så man kunne skrive direkte til memory location
		head := mz.Paths[0]

		head.X = 0

		println("dir: ", dir)
		fmt.Println("initial head", head)

		if dir == Left && head.X == 0 {
			continue
		}
		if dir == Right && head.X == xAxisBound-1 {
			continue
		}
		if dir == Up && head.Y == yAxisBound-1 {
			continue
		}
		if dir == Down && head.Y == 0 {
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

		mz.MazeTrack.Set(head.Y, head.X, 0)

		fmt.Println("head", head)

		mz.Paths[0] = head

		PrintMaze(mz)

	}

	return mz
}

func PrintMaze(mz Maze) {
	rows, cols := mz.MazeTrack.Dims()
	fmt.Print("  ", strings.Repeat("_ ", cols), "\n")
	for i := rows - 1; i >= 0; i-- {
		fmt.Print("| ")
		for j := 0; j < cols; j++ {
			fmt.Print(mz.MazeTrack.At(i, j), " ")
		}
		fmt.Print("|\n")
	}
	fmt.Print("  ", strings.Repeat("- ", cols), "\n")
}
