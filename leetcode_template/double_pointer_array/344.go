package double_pointer_array

func reverseString(s []byte) {
	for lo, hi := 0, len(s)-1; lo < hi; {
		s[lo], s[hi] = s[hi], s[lo]
		lo++
		hi--
	}
}
