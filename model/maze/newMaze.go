package maze

import (
	"fmt"
	"math/rand"

	"gonum.org/v1/gonum/mat"
)

func NewMaze(r, c int) Maze {
	data := make([]float64, r*c)
	for i := 0; i < r*c; i++ {
		data[i] = 1
	}

	mz := &Maze{
		Maze:   mat.NewDense(r, c, data),
		yBound: r,
		xBound: c,
	}

	createPath(mz)

	println(&mz.Maze)
	return *mz
}

func createPath(mz *Maze) {
	mz.Maze.Set(0, 0, float64(0))

	heads := make([]*coordinate, 0, 10)
	heads = append(heads, &coordinate{0, 0})
	var tmpHeads []*coordinate
	copy(tmpHeads, heads)
	for i := 0; i < 7; i++ {

		println("******************", i)
		for k := 0; k < len(heads); k++ {

			rows, cols := mz.Maze.Dims()
			// stepRipple(mz)

			convMat := findAvgConvMatrix(mz.Maze)
			fmt.Printf("Matrix:\n")
			matFormatted := mat.Formatted(convMat, mat.Prefix(""), mat.Squeeze())
			fmt.Printf("%.2f\n", matFormatted)
			target := newTarget(rows, cols)

			stepsPerRound := 4
			for j := 0; j < stepsPerRound; j++ {
				stepVectorProduct(mz, target, heads[k])
				PrintMaze(*mz, *heads[k])

				if heads[k].equals(target) {
					break
				}
				r := rand.Intn(stepsPerRound)
				if float64(r) <= 1.0/3*float64(stepsPerRound) {
					tmpHead := *heads[k]
					tmpHeads = append(tmpHeads, &tmpHead)
				}

			}
		}

		heads = tmpHeads
	}
}

func selectHead(heads []*coordinate) *coordinate {
	prob := float64(1) / float64(len(heads))
	summedProbs := make([]float64, len(heads))

	summedProbs[0] = prob
	head := rand.Float64()

	for i := 1; i < len(summedProbs); i++ {
		summedProbs[i] = summedProbs[i-1] + prob
	}

	for i := 0; i < len(summedProbs); i++ {
		if head < summedProbs[i] {
			return heads[i]
		}
	}
	return nil
}
