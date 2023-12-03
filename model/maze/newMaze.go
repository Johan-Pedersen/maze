package maze

import (
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

	heads := make([]*coordinate, 0, 10)
	heads = append(heads, &coordinate{0, 0})

	var curHead *coordinate
	for i := 0; i < 10; i++ {

		println("******************", i)

		rows, cols := mz.Maze.Dims()
		// stepRipple(mz)
		target := newTarget(rows, cols)

		curHead = selectHead(heads)

		if curHead == nil {
			log.Fatal("curHead = nil")
		}
		stepsPerRound := 5
		for j := 0; j < stepsPerRound; j++ {
			stepVectorProduct(mz, target, curHead)
			PrintMaze(*mz, *curHead)

			if curHead.equals(target) {
				break
			}
		}
		tmp_head := *curHead
		heads = append(heads, &tmp_head)

	}
}

func selectHead(heads []*coordinate) *coordinate {
	prob := float64(1) / float64(len(heads))
	summedProbs := make([]float64, len(heads))

	summedProbs[0] = prob
	head := rand.Float64()

	// Det her ville man kunne gøre meget smartere i scala
	for i := 1; i < len(summedProbs); i++ {
		summedProbs[i] = summedProbs[i-1] + prob
	}

	for i := 0; i < len(summedProbs); i++ {
		if head < summedProbs[i] {
			return heads[i]
		}
	}
	// Should be an error
	return nil
}
