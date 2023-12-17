package maze

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
)

func step(probs []float64, mz Maze, head *coordinate) {
	// Dette er selve step metoden
	dir, err := sample(probs)
	if err != nil {
		log.Fatal(err)
	}

	println("probs:")
	fmt.Println(probs)
	fmt.Println("\ndir:", dir)

	switch dir {
	case Left:
		head.x = head.x - 1
	case Right:
		head.x = head.x + 1
	case Up:
		head.y = head.y + 1
	case Down:
		head.y = head.y - 1
	}

	mz.Maze.Set(head.y, head.x, 0)
}

func sample(probs []float64) (stepDirection, error) {
	summedProbs := make([]float64, len(probs))

	summedProbs[0] = probs[0]
	step := rand.Float64()

	// Det her ville man kunne gøre meget smartere i scala
	for i := 1; i < len(probs); i++ {
		summedProbs[i] = summedProbs[i-1] + probs[i]
	}

	for i := 0; i < len(summedProbs); i++ {
		if step < summedProbs[i] {
			return stepDirection(i), nil
		}
	}
	return stepDirection(0), errors.New("could not sample directions. \nProbs: " + fmt.Sprintf("%v", probs) +
		"\nstep: " + strconv.FormatFloat(step, 'f', -1, 64))
}

func getValidDirs(mz *Maze, head coordinate) []stepDirection {
	dirs := make([]stepDirection, 0, 3)

	if isDirValid(Left, head, mz.xBound, mz.yBound) {
		if mz.Maze.At(head.y, head.x-1) != 0.00 {
			dirs = append(dirs, Left)
		}
	}
	if isDirValid(Right, head, mz.xBound, mz.yBound) {
		if mz.Maze.At(head.y, head.x+1) != 0.00 {
			dirs = append(dirs, Right)
		}
	}
	if isDirValid(Up, head, mz.xBound, mz.yBound) {
		if mz.Maze.At(head.y+1, head.x) != 0.00 {
			dirs = append(dirs, Up)
		}
	}
	if isDirValid(Down, head, mz.xBound, mz.yBound) {
		if mz.Maze.At(head.y-1, head.x) != 0.00 {
			dirs = append(dirs, Down)
		}
	}
	return dirs
}

/*
Checks if the step direction is valid from given path coordinate

*/

func isDirValid(dir stepDirection, head coordinate,
	xAxisBound, yAxisBound int,
) bool {
	return (dir == Left && head.x != 0) ||
		(dir == Right && head.x != xAxisBound-1) ||
		(dir == Up && head.y != yAxisBound-1) ||
		(dir == Down && head.y != 0)
}
