package maze

import (
	"gonum.org/v1/gonum/mat"
)

func ripple(mz *mat.Dense, target coordinate) {
	yBound, xBound := mz.Dims()
	for x := 0; x < xBound; x++ {
		for y := 0; y < yBound; y++ {
			// mz.Set(y, x, Norm(target.X-x, target.Y-y))
		}
	}
}

/*
Tag et step for en PathCoordinate

Hvad skal den gøre

  - Vi skal tage et step

  - Det skal vœre et valid step

  - Ud af de valid steps man kan tage vœlger man dem med en ss der går mod target zone.

  - Det betyder vi skal have have alle valid steps, deres totale sum, for hver dir skal vi så tillœkke en sandsynlig

  - head skal vœre en pointer, fordi så kan vi opdatere head direkte
*/
func stepRipple(mz *Maze) {
	// Valid dirs
	weights := make([]float64, 4)
	probs := make([]float64, 4)

	// Det er noget pointer shit.
	// Er det en pointer når det er et array ->
	head := &mz.paths[0]

	var distanceSum float64

	if isDirValid(Left, *head, mz.xBound, mz.yBound) {
		weights[Left] = mz.Maze.At(head.y, head.x-1)
		distanceSum += mz.Maze.At(head.y, head.x-1)
	}
	if isDirValid(Right, *head, mz.xBound, mz.yBound) {
		weights[Right] = mz.Maze.At(head.y, head.x+1)
		distanceSum += mz.Maze.At(head.y, head.x+1)
	}
	if isDirValid(Up, *head, mz.xBound, mz.yBound) {
		weights[Up] = mz.Maze.At(head.y+1, head.x)
		distanceSum += mz.Maze.At(head.y+1, head.x)
	}
	if isDirValid(Down, *head, mz.xBound, mz.yBound) {
		weights[Down] = mz.Maze.At(head.y-1, head.x)
		distanceSum += mz.Maze.At(head.y-1, head.x)
	}

	for i := 0; i < 4; i++ {
		// Dette er forkert, her er det den der ligger lœngest vœk som har størst ss for at blive valgt.
		probs[i] = weights[i] / distanceSum
	}

	step(probs, *mz)
}
