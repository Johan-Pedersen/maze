package maze

import (
	"fmt"
	"math"
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

	rn := rand.New(rand.NewSource(2))
	y = rn.Intn(yUpperBound-yLowerBound) + yLowerBound
	x = rn.Intn(xUpperBound-xLowerBound) + xLowerBound

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
	head := mz.Paths[0]
	fmt.Print("  ", strings.Repeat("_ ", cols), "\n")
	for i := rows - 1; i >= 0; i-- {
		fmt.Print("| ")
		for j := 0; j < cols; j++ {
			if head.X == j && head.Y == i {
				fmt.Printf(" X  ")
			} else {
				fmt.Printf("%.2f ", mz.Maze.At(i, j))
			}
		}
		fmt.Print("|\n")
	}
	fmt.Print("  ", strings.Repeat("- ", cols), "\n")
}

func Norm(x1, x2 int) float64 {
	return math.Sqrt(math.Pow(float64(x1), 2) + math.Pow(float64(x2), 2))
}
