package math

// GCD greatest common divisor via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM find Less Common Multiple via GCD
func LCM(first int, integers []int) int {
	result := first * integers[0] / GCD(first, integers[0])
	for i := 1; i < len(integers); i++ {
		result = LCM(result, []int{integers[i]})
	}
	return result
}
