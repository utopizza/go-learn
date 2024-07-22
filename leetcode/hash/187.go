package hash

func findRepeatedDnaSequences(s string) []string {
	resSet := make(map[string]bool)
	subSet := make(map[string]bool)

	for l, r := 0, 9; r < len(s); {
		sub := s[l : r+1]
		if subSet[sub] {
			resSet[sub] = true
		} else {
			subSet[sub] = true
		}
		l++
		r++
	}

	var res []string
	for k := range resSet {
		res = append(res, k)
	}

	return res
}
