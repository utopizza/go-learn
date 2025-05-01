package array

func plusOne(digits []int) []int {
	increment := 1
	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + increment
		digits[i] = sum % 10
		increment = sum / 10
	}
	if increment == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
