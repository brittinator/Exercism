package diffsquares

import "math"

const (
	testVersion = 1
)

func SquareOfSums(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum * sum
}

func SumOfSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

func Difference(n int) int {
	return int(math.Abs(float64(SumOfSquares(n)) - float64(SquareOfSums(n))))
}
