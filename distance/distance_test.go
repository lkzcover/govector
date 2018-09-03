package distance

import "testing"

func TestCosine(t *testing.T) {
	x := []float64{3, 4}
	y := []float64{3, 4}

	res, err := Cosine(x, y)
	if err != nil {
		t.Error(err.Error())
	}
	if res != 1 {
		t.Error("Expected 1, got", res)
	}
}
