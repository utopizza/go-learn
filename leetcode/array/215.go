package array

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	return quickSelect(nums, 0, n-1, n-k)
}

func quickSelect(nums []int, l, r, k int) int {
	if l == r {
		return nums[k]
	}
	partition := nums[l]
	i := l - 1
	j := r + 1
	for i < j {
		for i++; nums[i] < partition; i++ {
		}
		for j--; nums[j] > partition; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	if k <= j {
		return quickSelect(nums, l, j, k)
	} else {
		return quickSelect(nums, j+1, r, k)
	}
}
