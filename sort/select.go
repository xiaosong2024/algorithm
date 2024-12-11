package sort

// BubbleSort

// SelectSort 选择排序
// 选择排序
// 外层 i= [0~n-1]
// 内存 从 i~n-1 中选择最小的，和i交换
// 1w * 1k = 33.1962311秒
func SelectSort(arr []int) {
	// 选择最小的，放在第i个
	n := len(arr)
	for i := 0; i < n-1; i++ {
		min_idx, _ := findMin(arr, i, n-1)
		_swap(arr, i, min_idx)
	}
}
