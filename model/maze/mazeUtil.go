package maze

import (
	"fmt"
	"math/rand"
	"strings"
)

/*
Generate x,y targetzone
*/
func TargetZone(r, c int) (x, y int) {
	margin := 0.15

	yLowerBound := int(margin * float64(r))
	xLowerBound := int(margin * float64(c))

	yUpperBound := int((1 - margin) * float64(r))
	xUpperBound := int((1 - margin) * float64(c))

	y = rand.Intn(yUpperBound-yLowerBound) + yLowerBound
	x = rand.Intn(xUpperBound-xLowerBound) + xLowerBound

	return x, y
}

/*
Map between (x,y) and the 1 D representation for mat.NewDense data
*/
func CoordToDataPos(x, y, colDims int) int {
	return colDims*y + x
}

/*
Print maze in a nice way
*/
func PrintMaze(mz Maze) {
	rows, cols := mz.Maze.Dims()
	fmt.Print("  ", strings.Repeat("_ ", cols), "\n")
	for i := rows - 1; i >= 0; i-- {
		fmt.Print("| ")
		for j := 0; j < cols; j++ {
			fmt.Print(mz.Maze.At(i, j), " ")
		}
		fmt.Print("|\n")
	}
	fmt.Print("  ", strings.Repeat("- ", cols), "\n")
}
