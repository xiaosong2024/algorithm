package sort

// BubbleSort

// BubbleSort 冒泡排序
// 冒泡排序的思路是
// 内层循环：通过一次变量，将最小的从最后一位移动到第一位
// 外层循环：让从n=[0 ~ n-2]的每个值变成后续中最小的
// 10000条 * 1000 次 = 45.9216055秒
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
