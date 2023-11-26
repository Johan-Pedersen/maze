package maze

import (
	"fmt"
	"log"
	"math"
)

/*
Tag 1 skridt i en valid retning
*/
func stepVectorProduct(mz *Maze, target coordinate) {
	dirs := getValidDirs(mz)

	fmt.Println("Valid dirs", dirs)

	head := mz.paths[0]

	// find valid paths og vektorne til dem
	// Vektorne har vi allerede

	// find target koordinate og vector

	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	println("target: ", "(", target.x, target.y, ")")

	targetDir := newNormVector(head, target)

	fmt.Println("targetDir", targetDir.String())
	var productSum float64 = 0

	probs := make([]float64, 4)

	products := make([]float64, 4)

	for _, dir := range dirs {
		// find indre produkt
		switch dir {
		case Left:
			cellVector := newNormVector(head, coordinate{head.x - 1, head.y})
			fmt.Println("cell vector", cellVector.String())
			innerProd := innerProduct(cellVector, targetDir)
			productSum += max(0, float64(innerProd))
			fmt.Println("InnerProd: ", innerProd)
			fmt.Println("Angle between cellVector and target Vector:", math.Acos(innerProd)*180/math.Pi)
			fmt.Println("ProductSum: ", productSum)
			products[Left] = max(0, float64(innerProd))

		case Right:
			cellVector := newNormVector(head, coordinate{head.x + 1, head.y})

			fmt.Println("cell vector", cellVector.String())
			innerProd := innerProduct(cellVector, targetDir)
			productSum += max(0, float64(innerProd))
			fmt.Println("InnerProd: ", innerProd)
			fmt.Println("Angle between cellVector and target Vector:", math.Acos(innerProd)*180/math.Pi)
			fmt.Println("ProductSum: ", productSum)
			products[Right] = max(0, float64(innerProd))

		case Up:
			cellVector := newNormVector(head, coordinate{head.x, head.y + 1})
			fmt.Println("cell vector", cellVector.String())
			innerProd := innerProduct(cellVector, targetDir)
			productSum += max(0, float64(innerProd))
			fmt.Println("InnerProd: ", innerProd)
			fmt.Println("Angle between cellVector and target Vector:", math.Acos(innerProd)*180/math.Pi)
			fmt.Println("ProductSum: ", productSum)
			products[Up] = max(0, float64(innerProd))

		case Down:
			cellVector := newNormVector(head, coordinate{head.x, head.y - 1})
			fmt.Println("cell vector", cellVector.String())
			innerProd := innerProduct(cellVector, targetDir)
			productSum += max(0, float64(innerProd))
			fmt.Println("InnerProd: ", innerProd)
			fmt.Println("Angle between cellVector and target Vector:", math.Acos(innerProd)*180/math.Pi)
			fmt.Println("ProductSum: ", productSum)
			products[Down] = max(0, float64(innerProd))
		}
	}

	if productSum == 0 {
		log.Fatal("productSum = 0")
	}

	// find procent

	fmt.Println("")
	for i := 0; i < 4; i++ {
		// Dette er forkert, her er det den der ligger lœngest vœk som har størst ss for at blive valgt.
		probs[i] = products[i] / productSum
	}

	step(probs, *mz)

	// lav step mod target
	// stepToTarget(&head, maze.TargetCoordinate{tmpTargetx, tmpTargety})
}
