package array

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func reverse(x []int) {
	for lo, hi := 0, len(x)-1; lo < hi; {
		x[lo], x[hi] = x[hi], x[lo]
		lo++
		hi--
	}
}
