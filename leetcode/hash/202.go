package hash

func isHappy(n int) bool {
	set := make(map[int]bool)
	for {
		next := getNextHappyNum(n)
		if next == 1 {
			return true
		}
		if set[next] {
			return false
		}
		set[next] = true
		n = next
	}
}

func getNextHappyNum(x int) int {
	sum := 0
	for x != 0 {
		y := x % 10
		sum = sum + y*y
		x = x / 10
	}
	return sum
}
