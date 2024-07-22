package hash

func isIsomorphic(s string, t string) bool {
	return isIsomorphicCore(s, t) && isIsomorphicCore(t, s)
}

func isIsomorphicCore(s string, t string) bool {
	m := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if c, exist := m[s[i]]; exist {
			if c != t[i] {
				return false
			}
		} else {
			m[s[i]] = t[i]
		}
	}
	return true
}
