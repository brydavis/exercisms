package diffsquares

import "math"

func SquareOfSums(n int) int {
	var s int
	for i := 0; i <= n; i++ {
		s += i
	}
	return int(math.Pow(float64(s), 2.0))
}

func SumOfSquares(n int) int {
	var s float64
	for i := 0; i <= n; i++ {
		s += math.Pow(float64(i), 2.0)
	}
	return int(s)
}

func Difference(n int) int {
	sumSqr := SumOfSquares(n)
	sqrSum := SquareOfSums(n)
	return sqrSum - sumSqr
}
