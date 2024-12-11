package sort

// BubbleSort

// QuickSort 快速排序
// 0.7 秒
func QuickSort(arr []int) {
	_quickSort(arr)
}

func _quickSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}
	l, r := partition(arr)
	// 左边有没有元素
	_quickSort(arr[0:l])
	_quickSort(arr[r+1 : n])
}

// 返回这个数组中 等于第一项的开始，结束位置
func partition(arr []int) (int, int) {
	var (
		// 定义 ：
		// 小于 0~i
		// 等于 i+1 ~ j
		// 大于 k ~ n-1
		n   = len(arr)
		i   = -1
		j   = 0
		k   = n
		e   = 1 // 当前遍历指针
		val = arr[0]
	)

	//utils.VarDump(arr)
	//utils.VarDump("val=", val)
	//utils.VarDump("arr[e]=", arr[e])
	for e <= k-1 {
		//utils.VarDump(arr, e)
		if arr[e] < val {
			_swap(arr, e, i+1)
			i++
			j++
			e++
		} else if arr[e] == val {
			j++
			e++
		} else {
			_swap(arr, e, k-1)
			k--
		}
	}
	return i + 1, j
}
