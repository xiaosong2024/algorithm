package sort

// InsertSort 插入排序
// 第一个元素已经是有序的了
// 遍历从 i=[1~n-1] 将每个元素移动到 j=[i-1 ~ 0]的合适位置
// 外层 i=[1 ~ n-1]
// 内层 j=[i-1 ~ 0] 如果 arr[j+1] < arr[j] 则交换，否则 j 元素已经在合适的位置了 退出
// 1w*1k=33.4339259s
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
