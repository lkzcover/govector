package basic

import "testing"

func TestSProduct(t *testing.T) {
	x := []float64{1, 2, 3}
	y := []float64{1, 2, 3}

	res, err := SProduct(x, y)
	if err != nil {
		t.Error(err.Error())
	}
	if res != 14 {
		t.Error("Expected 14, got", res)
	}
}
