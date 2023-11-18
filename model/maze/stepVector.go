package maze

/*
Tag 1 skridt i en valid retning
*/
func stepVectorProduct(mz *Maze) {
	dirs := validDirs(mz)

	head := mz.Paths[0]

	// find valid paths og vektorne til dem
	// Vektorne har vi allerede

	// find target koordinate og vector
	rows, cols := mz.Maze.Dims()

	target := GenTargetZone(rows, cols)

	targetDir := NewNormVector(head, target)

	var productSum float64 = 0

	probs := make([]float64, 4)

	products := make([]float64, 4)

	for _, dir := range dirs {
		// find indre produkt
		switch dir {
		case Left:
			cellVector := NewNormVector(head, coordinate{head.X - 1, head.Y})
			innerProd := InnerProduct(cellVector, targetDir)

			productSum += float64(innerProd)
			products[Left] = float64(innerProd)

		case Right:
			cellVector := NewNormVector(head, coordinate{head.X + 1, head.Y})
			innerProd := InnerProduct(cellVector, targetDir)
			productSum += float64(innerProd)
			products[Right] = float64(innerProd)
		case Up:
			cellVector := NewNormVector(head, coordinate{head.X, head.Y + 1})
			innerProd := InnerProduct(cellVector, targetDir)
			productSum += float64(innerProd)

			products[Up] = float64(innerProd)
		case Down:
			cellVector := NewNormVector(head, coordinate{head.X, head.Y - 1})
			innerProd := InnerProduct(cellVector, targetDir)
			productSum += float64(innerProd)

			products[Down] = float64(innerProd)
		}
	}

	// find procent

	for i := 0; i < 4; i++ {
		// Dette er forkert, her er det den der ligger lœngest vœk som har størst ss for at blive valgt.
		probs[i] = products[i] / productSum
	}

	step(probs, *mz)

	// lav step mod target
	// stepToTarget(&head, maze.TargetCoordinate{tmpTargetX, tmpTargetY})
}
