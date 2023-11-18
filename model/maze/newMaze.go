package maze

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func NewMaze(r, c int) Maze {
	data := make([]float64, r*c)
	for i := 0; i < r*c; i++ {
		data[i] = 1
	}

	// Generate 2x2 target zone
	mz := &Maze{
		Maze:   mat.NewDense(r, c, data),
		Target: GenTargetZone(r, c),
		Paths:  []Coordinate{},
		YBound: r,
		XBound: c,
	}

	data2 := make([]float64, r*c)

	// ripple(mz.Maze, mz.Target)

	copy(data2, data)

	maze2 := mat.NewDense(r, c, data2)

	// ripple(maze2, mz.Target)

	mz2 := Maze{
		Maze:   maze2,
		Target: mz.Target,
		Paths:  append(mz.Paths, Coordinate{0, 0}),
		YBound: mz.YBound,
		XBound: mz.XBound,
	}

	createPath(mz)
	println("Target: (", mz.Target.X, mz.Target.Y, ")")

	PrintMaze(mz2)
	println(&maze2)
	println(&mz.Maze)
	return *mz
}

func createPath(mz *Maze) {
	// Init first position
	mz.Paths = append(mz.Paths, Coordinate{0, 0})

	// mzTrack := mz.Maze
	mz.Maze.Set(0, 0, float64(0))

	// direction := [4] StepDirection {Left, Right, Up, Down}

	println("left:", Left)
	println("Right", Right)
	println("Up:", Up)
	println("Down:", Down)
	fmt.Println("dims", mz.XBound, mz.YBound)
	// Tager 10 skridt
	for i := 0; i < 20; i++ {

		println("******************", i)

		// stepRipple(mz)

		StepVectorProduct(mz)

		PrintMaze(*mz)

	}
}
