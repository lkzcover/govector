package basic

import (
	"errors"
)

//SProduct scalar product 2 vector
func SProduct(x, y []float64) (res float64, err error) {
	if len(x) != len(y) {
		return 0, errors.New("Vector len incorrect")
	}

	for i, x1 := range x {
		res = res + (x1 * y[i])
	}

	return res, nil
}
