package sort

func _swap(arr []int, i, j int) {
	if i == j {
		return
	}
	arr[i], arr[j] = arr[j], arr[i]
}

func findMin(arr []int, start, end int) (int, int) {
	// 返回index和value
	min_idx := start
	min_value := arr[start]
	for i := start; i <= end; i++ {
		if arr[i] < min_value {
			min_idx = i
			min_value = arr[i]
		}
	}
	return min_idx, min_value
}
