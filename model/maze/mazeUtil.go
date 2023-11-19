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

func GenTargetZone(r, c int) coordinate {
	margin := 0.15

	yLowerBound := int(margin * float64(r))
	xLowerBound := int(margin * float64(c))

	yUpperBound := int((1 - margin) * float64(r))
	xUpperBound := int((1 - margin) * float64(c))

	// rn := rand.New(rand.NewSource(2))
	y := rand.Intn(yUpperBound-yLowerBound) + yLowerBound
	x := rand.Intn(xUpperBound-xLowerBound) + xLowerBound

	return coordinate{x, y}
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
		fmt.Print(i, "| ")
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
	fmt.Print("   ")
	for i := 0.0; i < float64(cols); i++ {
		fmt.Printf("%.2f ", i)
	}
	fmt.Println()
}

/*
Norm of vector vec (x y)
*/
func Norm(vec Vector) float64 {
	return math.Sqrt(math.Pow(float64(vec.X), 2) + math.Pow(float64(vec.Y), 2))
}

/*
input vectors: (x1,x2), (y1,y2)
*/
func InnerProduct(vec1 Vector, vec2 Vector) float64 {
	return vec1.X*vec2.X + vec1.Y*vec2.Y
}
