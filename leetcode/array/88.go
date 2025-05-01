package array

func merge(nums1 []int, m int, nums2 []int, n int) {
	temp := make([]int, m)
	for i := 0; i < m; i++ {
		temp[i] = nums1[i]
	}
	for i, j, k := 0, 0, 0; i < m || j < n; k++ {
		if i == m {
			nums1[k] = nums2[j]
			j++
		} else if j == n {
			nums1[k] = temp[i]
			i++
		} else if temp[i] <= nums2[j] {
			nums1[k] = temp[i]
			i++
		} else {
			nums1[k] = nums2[j]
			j++
		}
	}
}
