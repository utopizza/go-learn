package array

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	for start := 0; start < n; {
		sum := 0
		step := 0
		for step < n {
			location := (start + step) % n
			sum += gas[location]
			sum -= cost[location]
			if sum < 0 {
				break
			}
			step++
		}
		if step == n {
			return start
		} else {
			start = start + step + 1
		}
	}
	return -1
}
