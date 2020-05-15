package diffsquares

//The way a computer scienties will solve the problem.

// SquareOfSum returns a sum of square of given num of integers
func SquareOfSum(n int) int {
	return sumOfIntergers(n) * sumOfIntergers(n)
}

func sumOfIntergers(n int) int {
	if n <= 0 {
		return n
	} else {
		return n + sumOfIntergers(n-1)
	}
}

// SumOfSquares returns the sum of squares of given integer
func SumOfSquares(n int) int {
	if n <= 1 {
		return n
	} else {
		return n*n + SumOfSquares(n-1)
	}
}

// Difference returns the differnce
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

/**
// The way a physicist or mathematician will solve this problem
// SquareOfSum returns a sum of square of given num of integers
func SquareOfSum(n int) int {
	sumInt := n * (n + 1) / 2
	return sumInt * sumInt
}

// SumOfSquares -
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference -
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
**/
