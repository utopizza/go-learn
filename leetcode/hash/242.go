package hash

func isAnagram(s string, t string) bool {
	sMap := buildCountMap(s)
	tMap := buildCountMap(t)
	return isSameMap(sMap, tMap)
}

func buildCountMap(s string) map[byte]int {
	m := make(map[byte]int)
	for i := range s {
		m[s[i]]++
	}
	return m
}

func isSameMap(m1, m2 map[byte]int) bool {
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	for k, v := range m2 {
		if m1[k] != v {
			return false
		}
	}
	return true
}
