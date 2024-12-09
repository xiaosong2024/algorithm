package sort

func InsertSort(arr []int) {
	n := len(arr)
	for i := 1; i <= n-1; i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[j+1] < arr[j] {
				_swap(arr, j, j+1)
			} else {
				break
			}
		}
	}
}
