package maze

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

/*
Find the avg Convolution matrix

Applies the avg kernel, on the maze matrix.

Kernel size 1/3 of original size
*/
func findAvgConvMatrix(maze *mat.Dense) *mat.Dense {
	rows, cols := maze.Dims()

	downScale := 3
	scaledRows := float64(rows) / float64(downScale)
	scaledCols := float64(cols) / float64(downScale)

	data := make([]float64, int(scaledRows*scaledCols))
	for i := 0; i < int(scaledRows*scaledCols); i++ {
		data[i] = 1
	}

	convMat := mat.NewDense(int(scaledRows), int(scaledCols), data)

	for i := 0; i < int(scaledRows); i++ {
		for j := 0; j < int(scaledCols); j++ {
			res := 0.0
			// rows -> y-akse
			for p := 0; p < downScale; p++ {
				// cols -> x-akse
				for q := 0; q < downScale; q++ {
					res += maze.At(downScale*i+p, downScale*j+q)
				}
			}
			convMat.Set(i, j, res/math.Pow(float64(downScale), 2))
		}
	}

	return convMat
}
