package distance

import (
	"errors"

	"github.com/lkzcover/govector/basic"
	"github.com/lkzcover/govector/norma"
)

//Cosine cosine distance for 2 vector
func Cosine(x, y []float64) (res float64, err error) {
	if len(x) != len(y) {
		return 0, errors.New("Vector len incorrect")
	}

	sp, _ := basic.SProduct(x, y)
	return sp / (norma.L2Norm(x) * norma.L2Norm(y)), nil
}
