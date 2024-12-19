package numeric

import "math"

func NDigits(v float64) int {
	return int(math.Log10(v) + 1)
}

// Splits a number in half based on its digits.
// ie 2024 become 20, 24
func SplitNum(n int) (int, int) {
	numDigits := int(math.Log10(float64(n))) + 1
	halfDigits := numDigits / 2

	// Calculate the divisor to split the number
	divisor := int(math.Pow(10, float64(halfDigits)))

	// Split the number
	right := n % divisor
	left := n / divisor

	return left, right
}
