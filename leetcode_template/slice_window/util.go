package slice_window

func isSameSet(set1, set2 map[byte]int) bool {
	for k, v := range set1 {
		if set2[k] != v {
			return false
		}
	}
	for k, v := range set2 {
		if set1[k] != v {
			return false
		}
	}
	return true
}

func isContainsSet(set1, set2 map[byte]int) bool {
	for k, v := range set2 {
		if set1[k] < v {
			return false
		}
	}
	return true
}
