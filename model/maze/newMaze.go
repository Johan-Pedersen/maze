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

	mz := &Maze{
		Maze:   mat.NewDense(r, c, data),
		Target: genTargetZone(r, c),
		yBound: r,
		xBound: c,
	}

	createPath(mz)
	println("Target: (", mz.Target.x, mz.Target.y, ")")

	println(&mz.Maze)
	return *mz
}

func createPath(mz *Maze) {
	mz.Maze.Set(0, 0, float64(0))

	println("left:", Left)
	println("Right", Right)
	println("Up:", Up)
	println("Down:", Down)
	fmt.Println("dims", mz.xBound, mz.yBound)
	// Tager 10 skridt
	head := coordinate{0, 0}
	for i := 0; i < 3; i++ {

		println("******************", i)

		rows, cols := mz.Maze.Dims()
		// stepRipple(mz)
		target := genTargetZone(rows, cols)

		for j := 0; j < 5; j++ {
			stepVectorProduct(mz, target, &head)
			PrintMaze(*mz, head)

			/*
				if mz.paths[0].equals(target) {
					break
				}
			*/
		}

	}
}
