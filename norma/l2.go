package norma

import "math"

//L2Norm - Euclidean norm
func L2Norm(x []float64) (res float64) {
	for _, xi := range x {
		res = res + math.Pow(xi, 2)
	}
	return math.Sqrt(res)
}
