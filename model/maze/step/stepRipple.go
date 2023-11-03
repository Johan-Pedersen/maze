package step

import "gonum.org/v1/gonum/mat"

func ripple(maze *mat.Dense, target TargetCoordinate) {
	yBound, xBound := maze.Dims()
	for x := 0; x < xBound; x++ {
		for y := 0; y < yBound; y++ {
			maze.Set(y, x, Norm(target.X-x, target.Y-y))
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
func stepRipple(maze *Maze) {
	// Valid dirs
	weights := make([]float64, 4)
	probs := make([]float64, 4)

	// Det er noget pointer shit.
	// Er det en pointer når det er et array ->
	head := &maze.Paths[0]

	var distanceSum float64

	if isDirValid(Left, *head, maze.XBound, maze.YBound) {
		weights[Left] = maze.Maze.At(head.Y, head.X-1)
		distanceSum += maze.Maze.At(head.Y, head.X-1)
	}
	if isDirValid(Right, *head, maze.XBound, maze.YBound) {
		weights[Right] = maze.Maze.At(head.Y, head.X+1)
		distanceSum += maze.Maze.At(head.Y, head.X+1)
	}
	if isDirValid(Up, *head, maze.XBound, maze.YBound) {
		weights[Up] = maze.Maze.At(head.Y+1, head.X)
		distanceSum += maze.Maze.At(head.Y+1, head.X)
	}
	if isDirValid(Down, *head, maze.XBound, maze.YBound) {
		weights[Down] = maze.Maze.At(head.Y-1, head.X)
		distanceSum += maze.Maze.At(head.Y-1, head.X)
	}

	for i := 0; i < 4; i++ {
		// Dette er forkert, her er det den der ligger lœngest vœk som har størst ss for at blive valgt.
		probs[i] = weights[i] / distanceSum
	}

	step(probs, maze)
}
