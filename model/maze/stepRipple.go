package maze

import (
	"gonum.org/v1/gonum/mat"
)

func ripple(mz *mat.Dense, target Coordinate) {
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
	head := &mz.Paths[0]

	var distanceSum float64

	if isDirValid(Left, *head, mz.XBound, mz.YBound) {
		weights[Left] = mz.Maze.At(head.Y, head.X-1)
		distanceSum += mz.Maze.At(head.Y, head.X-1)
	}
	if isDirValid(Right, *head, mz.XBound, mz.YBound) {
		weights[Right] = mz.Maze.At(head.Y, head.X+1)
		distanceSum += mz.Maze.At(head.Y, head.X+1)
	}
	if isDirValid(Up, *head, mz.XBound, mz.YBound) {
		weights[Up] = mz.Maze.At(head.Y+1, head.X)
		distanceSum += mz.Maze.At(head.Y+1, head.X)
	}
	if isDirValid(Down, *head, mz.XBound, mz.YBound) {
		weights[Down] = mz.Maze.At(head.Y-1, head.X)
		distanceSum += mz.Maze.At(head.Y-1, head.X)
	}

	for i := 0; i < 4; i++ {
		// Dette er forkert, her er det den der ligger lœngest vœk som har størst ss for at blive valgt.
		probs[i] = weights[i] / distanceSum
	}

	step(probs, *mz)
}
