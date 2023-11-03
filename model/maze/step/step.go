package step

import (
	"fmt"
	"math/rand"

	"maze/model/maze"
)

func step(probs []float64, mz *maze.Maze) {
	// Dette er selve step metoden
	dir := Sample(probs)

	head := &mz.Paths[0]

	println("probs:")
	println("Left:", maze.Left, "Right:", maze.Right, "Up:", maze.Up, "Down:", maze.Down)
	for _, v := range probs {
		fmt.Print(v, " ")
	}
	println("\ndir:", dir)

	switch dir {
	case maze.Left:
		head.X = head.X - 1
	case maze.Right:
		head.X = head.X + 1
	case maze.Up:
		head.Y = head.Y + 1
	case maze.Down:
		head.Y = head.Y - 1
	}

	mz.Maze.Set(head.Y, head.X, 0)
}

func Sample(probs []float64) maze.StepDirection {
	summedProbs := make([]float64, len(probs))

	// Det her er vel noget pointer shit

	// Når man arbejder med arrays, så er det vel pointers til arrayet i stedet for

	// Undersøg
	summedProbs[0] = probs[0]
	step := rand.Float64()

	// Det her ville man kunne gøre meget smartere i scala
	for i := 1; i < len(probs); i++ {
		summedProbs[i] = summedProbs[i-1] + probs[i]
	}

	for i := 0; i < len(summedProbs); i++ {
		if step < summedProbs[i] {
			return maze.StepDirection(i)
		}
	}
	// Should be an error
	return maze.StepDirection(0)
}

func validDirs(mz *maze.Maze) []maze.StepDirection {
	dirs := make([]maze.StepDirection, 3)

	head := mz.Paths[0]

	if isDirValid(maze.Left, head, mz.XBound, mz.YBound) {
		dirs = append(dirs, maze.Left)
	}
	if isDirValid(maze.Right, head, mz.XBound, mz.YBound) {
		dirs = append(dirs, maze.Right)
	}
	if isDirValid(maze.Up, head, mz.XBound, mz.YBound) {
		dirs = append(dirs, maze.Up)
	}
	if isDirValid(maze.Down, head, mz.XBound, mz.YBound) {
		dirs = append(dirs, maze.Down)
	}
	return dirs
}

/*
Checks if the step direction is valid from given path coordinate

Kan fejle ved dir == Down og head.Y = 0
*/

func isDirValid(dir maze.StepDirection, head maze.PathCoordinate,
	xAxisBound, yAxisBound int,
) bool {
	return (dir == maze.Left && head.X != 0) ||
		(dir == maze.Right && head.X != xAxisBound-1) ||
		(dir == maze.Up && head.Y != yAxisBound-1) ||
		(dir == maze.Down && head.Y != 0)
}
