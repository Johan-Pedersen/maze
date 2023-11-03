package step

import "maze/model/maze"

/*
Rename, så det ikke hedder noget med step
*/
func stepVectorProduct(mz *maze.Maze) {
	dirs := validDirs(mz)

	head := mz.Paths[0]

	// find valid paths og vektorne til dem
	// Vektorne har vi allerede

	// find target koordinate og vector
	rows, cols := mz.Maze.Dims()

	tempTargetX, tempTargetY := maze.TargetZone(rows, cols)

	for _, dir := range dirs {
		// find indre produkt
		switch dir {
		case maze.Left:
			head.X = head.X - 1
			innerProduct := maze.InnerProduct()
		case maze.Right:
			head.X = head.X + 1
		case maze.Up:
			head.Y = head.Y + 1
		case maze.Down:
			head.Y = head.Y - 1
		}
	}

	// find procent

	// lav step mod target
	stepToTarget(&head, maze.TargetCoordinate{tempTargetX, tempTargetY})
}

/*
Laver update til direkte på memory location
*/
func stepToTarget(head *maze.PathCoordinate, target maze.TargetCoordinate) {
}
