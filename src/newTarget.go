package maze

import (
	"errors"
	"fmt"
	"math/rand"

	"gonum.org/v1/gonum/mat"
)

/*
Target coordinate
*/
func newTarget(convMat *mat.Dense, head coordinate) (coordinate, error) {
	r, c := convMat.Dims()
	searchDistBound := 7.0
	coords := make([]coordinate, 0, r*c)

	countCoords := 0
	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {

			tmpVec := newVector(head, coordinate{i*3 + 1, j*3 + 1})

			norm := norm(tmpVec)
			if norm < searchDistBound && convMat.At(j, i) > 0.6 {

				coords = append(coords, coordinate{i*3 + 1, j*3 + 1})
				countCoords++
			}
		}
	}

	if countCoords == 0 {
		fmt.Print("countCoords ==", countCoords)

		return coordinate{0, 0}, errors.New("no valid coords")
	}
	return coords[rand.Intn(countCoords)], nil
}
