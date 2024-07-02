package maze

import (
	"fmt"
	"log"
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

	heads := make([]*coordinate, 0, 7)
	// heads = append(heads, &coordinate{0, 0}, &coordinate{14, 14}, &coordinate{0, 14}, &coordinate{14, 0}, &coordinate{7, 14}, &coordinate{7, 0})
	heads = append(heads, &coordinate{0, 0}, &coordinate{14, 14}, &coordinate{0, 14}, &coordinate{14, 0}, &coordinate{7, 7})
	// var tmpHeads []*coordinate
	// copy(tmpHeads, heads)
	for i := 0; i < 7; i++ {

		println("******************", i)
		for k := 0; k < len(heads); k++ {

			convMat := findAvgConvMatrix(mz.Maze)
			fmt.Printf("Convolution Matrix:\n")
			matFormatted := mat.Formatted(convMat, mat.Prefix(""), mat.Squeeze())
			fmt.Printf("%.2f\n", matFormatted)

			target, err := newTarget(convMat, *heads[k])

			if err != nil {
				log.Println(err)
			} else {

				stepsPerRound := 4
				for j := 0; j < stepsPerRound; j++ {
					prevHead := *heads[k]

					if heads[k].equals(target) {
						break
					}

					stepVectorProduct(mz, target, heads[k])

					if prevHead.equals(*heads[k]) {
						break
					}
					PrintMaze(*mz, *heads[k])

					if len(heads) < cap(heads) {

						r := rand.Intn(100)
						if r <= 1 {
							tmpHead := *heads[k]
							heads = append(heads, &tmpHead)
							// tmpHeads = append(tmpHeads, &tmpHead)
						}
					}

				}
			}
			// heads = tmpHeads
		}
	}

	fmt.Println("Heads")
	fmt.Println(heads)
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
