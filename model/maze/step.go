package maze

import (
	"fmt"
	"math/rand"
)

func step(probs []float64, mz Maze) {
	// Dette er selve step metoden
	dir := sample(probs)

	head := &mz.Paths[0]

	println("probs:")
	println("Left:", Left, "Right:", Right, "Up:", Up, "Down:", Down)
	for _, v := range probs {
		fmt.Print(v, " ")
	}
	println("\ndir:", dir)

	switch dir {
	case Left:
		head.X = head.X - 1
	case Right:
		head.X = head.X + 1
	case Up:
		head.Y = head.Y + 1
	case Down:
		head.Y = head.Y - 1
	}

	mz.Maze.Set(head.Y, head.X, 0)
}

func sample(probs []float64) StepDirection {
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
			return StepDirection(i)
		}
	}
	// Should be an error
	return StepDirection(0)
}

func validDirs(mz *Maze) []StepDirection {
	dirs := make([]StepDirection, 0, 3)

	head := mz.Paths[0]

	if isDirValid(Left, head, mz.XBound, mz.YBound, mz) {
		dirs = append(dirs, Left)
	}
	if isDirValid(Right, head, mz.XBound, mz.YBound, mz) {
		dirs = append(dirs, Right)
	}
	if isDirValid(Up, head, mz.XBound, mz.YBound, mz) {
		dirs = append(dirs, Up)
	}
	if isDirValid(Down, head, mz.XBound, mz.YBound, mz) {
		dirs = append(dirs, Down)
	}
	return dirs
}

/*
Checks if the step direction is valid from given path coordinate

Kan fejle ved dir == Down og head.Y = 0
*/

func isDirValid(dir StepDirection, head coordinate,
	xAxisBound, yAxisBound int, mz *Maze,
) bool {
	return (dir == Left && head.X != 0 && mz.Maze.At(head.X-1, head.Y) != 0) ||
		(dir == Right && head.X != xAxisBound-1 && mz.Maze.At(head.X+1, head.Y) != 0) ||
		(dir == Up && head.Y != yAxisBound-1 && mz.Maze.At(head.X, head.Y+1) != 0) ||
		(dir == Down && head.Y != 0 && mz.Maze.At(head.X, head.Y-1) != 0)
}
