package src

// Add returns the sum of two integers.
func Add(a, b int) int {
	return a + b
}

// Factorial returns the factorial of n.
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}
