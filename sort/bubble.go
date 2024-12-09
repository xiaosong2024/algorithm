package sort

// BubbleSort

// BubbleSort 冒泡排序
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := n - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				_swap(arr, j, j-1)
			}
		}
	}
}
