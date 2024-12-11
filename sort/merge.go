package sort

// BubbleSort

// MergeSort 归并排序
// 归并排序的思路
// 1w*1k=1.3056852
func MergeSort(arr []int) {
	res := _mergeSort(arr)
	for k, v := range res {
		arr[k] = v
	}
}

func _mergeSort(arr []int) []int {
	n := len(arr)
	if n == 1 {
		return arr
	}

	mid := n / 2
	left_res := _mergeSort(arr[0:mid])
	right_res := _mergeSort(arr[mid:n])
	return _merge(left_res, right_res)
}

// 关键函数：将两个有序数组合并，合并之后的数组仍然有序
func _merge(a, b []int) []int {
	var (
		a_len               = len(a)
		b_len               = len(b)
		c_len               = a_len + b_len
		c                   = make([]int, c_len)
		a_idx, b_idx, c_idx = 0, 0, 0
	)
	for ; c_idx <= c_len-1; c_idx++ {
		if a_idx == a_len {
			c[c_idx] = b[b_idx]
			b_idx++
		} else if b_idx == b_len {
			c[c_idx] = a[a_idx]
			a_idx++
		} else if a[a_idx] < b[b_idx] {
			c[c_idx] = a[a_idx]
			a_idx++
		} else {
			c[c_idx] = b[b_idx]
			b_idx++
		}
	}
	return c
}
