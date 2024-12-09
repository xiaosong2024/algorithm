package sort

// BubbleSort

// SelectSort 选择排序
func SelectSort(arr []int) {
	// 选择最小的，放在第i个
	n := len(arr)
	for i := 0; i < n-1; i++ {
		min_idx, _ := findMin(arr, i, n-1)
		_swap(arr, i, min_idx)
	}
}
