package hash

func groupAnagrams(strs []string) [][]string {
	mp := map[[26]int][]string{} // 骚，直接拿计数的数组作为key；如果语言不支持，需要先拼一个string
	for _, str := range strs {
		cnt := [26]int{}
		for _, b := range str {
			cnt[b-'a']++
		}
		mp[cnt] = append(mp[cnt], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
