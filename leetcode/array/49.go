package array

import "sort"

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		k := groupAnagramsGetStrKey(str)
		m[k] = append(m[k], str)
	}
	var result [][]string
	for _, group := range m {
		result = append(result, group)
	}
	return result
}

func groupAnagramsGetStrKey(str string) string {
	arr := []rune(str)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] <= arr[j]
	})
	return string(arr)
}
