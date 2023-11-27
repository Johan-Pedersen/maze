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

func genTargetZone(r, c int) coordinate {
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
func coordToDataPos(x, y, colDims int) int {
	return colDims*y + x
}

/*
Print maze in a nice way
*/
func PrintMaze(mz Maze) {
	rows, cols := mz.Maze.Dims()
	head := mz.paths[0]
	fmt.Print("  ", strings.Repeat("_ ", cols), "\n")
	for i := rows - 1; i >= 0; i-- {
		if i < 10 {
			fmt.Print(i, " | ")
		} else {
			fmt.Print(i, "| ")
		}
		for j := 0; j < cols; j++ {
			if head.x == j && head.y == i {
				fmt.Printf(" X   ")
			} else {
				cell := mz.Maze.At(i, j)
				if cell == 0 {
					fmt.Printf("\033[31m"+"%.2f ", mz.Maze.At(i, j))
				} else {
					fmt.Printf("\033[37m"+"%.2f ", mz.Maze.At(i, j))
				}
			}
		}
		fmt.Print("|\n")
	}
	fmt.Print("   ", strings.Repeat("- ", cols), "\n")
	fmt.Print("     ")
	for i := 0; i < cols; i++ {
			fmt.Printf("%-5d", i)
	}
	fmt.Println()
}

/*
Norm of vector vec (x y)
*/
func norm(vec vector) float64 {
	return math.Sqrt(math.Pow(float64(vec.x), 2) + math.Pow(float64(vec.y), 2))
}

/*
input vectors: (x1,x2), (y1,y2)
*/
func innerProduct(vec1 vector, vec2 vector) float64 {
	return vec1.x*vec2.x + vec1.y*vec2.y
}
