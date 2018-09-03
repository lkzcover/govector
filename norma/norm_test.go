package norma

import "testing"

func TestL2Norm(t *testing.T) {

	x := make([]float64, 2)

	x[0] = 3
	x[1] = 4

	if y := L2Norm(x); y != 5 {
		t.Error("Expected 5, got ", y)
	}

}
