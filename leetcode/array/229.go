package array

func majorityElementV2(nums []int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	var res []int
	for num, cnt := range m {
		if cnt > len(nums)/3 {
			res = append(res, num)
		}
	}
	return res
}
