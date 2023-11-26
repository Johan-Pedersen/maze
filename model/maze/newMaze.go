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
		Target: genTargetZone(r, c),
		paths:  []coordinate{},
		yBound: r,
		xBound: c,
	}

	data2 := make([]float64, r*c)

	// ripple(mz.Maze, mz.Target)

	copy(data2, data)

	maze2 := mat.NewDense(r, c, data2)

	// ripple(maze2, mz.Target)

	mz2 := Maze{
		Maze:   maze2,
		Target: mz.Target,
		paths:  append(mz.paths, coordinate{0, 0}),
		yBound: mz.yBound,
		xBound: mz.xBound,
	}

	createPath(mz)
	println("Target: (", mz.Target.x, mz.Target.y, ")")

	PrintMaze(mz2)
	println(&maze2)
	println(&mz.Maze)
	return *mz
}

func createPath(mz *Maze) {
	// Init first position
	mz.paths = append(mz.paths, coordinate{0, 0})

	// mzTrack := mz.Maze
	mz.Maze.Set(0, 0, float64(0))

	// direction := [4] StepDirection {Left, Right, Up, Down}

	println("left:", Left)
	println("Right", Right)
	println("Up:", Up)
	println("Down:", Down)
	fmt.Println("dims", mz.xBound, mz.yBound)
	// Tager 10 skridt
	for i := 0; i < 3; i++ {

		println("******************", i)

		rows, cols := mz.Maze.Dims()
		// stepRipple(mz)
		target := genTargetZone(rows, cols)

		for j := 0; j < 5; j++ {
			stepVectorProduct(mz, target)
			PrintMaze(*mz)

			if mz.paths[0].equals(target) {
				break
			}
		}

	}
}
