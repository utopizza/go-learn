package sorting

// InsertionSort ζε₯ζεΊ
func InsertionSort(arr []int) {
	for i := range arr {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			swap(arr, j, j-1)
		}
	}
}
