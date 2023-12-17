package maze

import (
	"fmt"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestAvgKernelDownScale(t *testing.T) {
	r, c := 15, 15
	data := make([]float64, r*c)
	for i := 0; i < r*c; i++ {
		data[i] = 1
	}

	mz := mat.NewDense(r, c, data)

	convMat := findAvgConvMatrix(mz)

	convR, convC := convMat.Dims()
	if float64(convR) != float64(r)/3.0 || float64(convC) != float64(c)/3.0 {
		t.Error(convR, "!=", float64(r)/3.0, "||", convC, "!=", float64(c)/3.0)
	}
}

func TestAvgKernelAllEntries1(t *testing.T) {
	r, c := 15, 15
	data := make([]float64, r*c)
	for i := 0; i < r*c; i++ {
		data[i] = 1
	}

	mz := mat.NewDense(r, c, data)

	convMat := findAvgConvMatrix(mz)

	fmt.Println(convMat.RawMatrix().Data)
	for _, v := range convMat.RawMatrix().Data {
		if v != 1 {
			t.Error(v, "!= 1")
		}
	}
}
