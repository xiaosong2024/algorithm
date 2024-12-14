package sort

// BubbleSort

// HeapSort 堆排序
func HeapSort(arr []int) {

	n := len(arr)
	heap := make([]int, n+1)
	for idx, val := range arr {
		heap[idx+1] = val
		_shiftUp(heap, idx+1)
	}
	// 从heap的第一个元素开始变量
	for i := 1; i <= n; i++ {
		arr[i-1] = heap[1] // 取堆顶的元素，然后shiftdown
		heap[1] = heap[n-i+1]
		_shiftDown(heap, 1)
	}
}

func _shiftUp(heap []int, idx int) {
	// 如果比自己的父亲小，则交换
	if idx == 1 {
		return
	}
	father_idx := idx / 2
	if heap[idx] < heap[father_idx] {
		_swap(heap, idx, father_idx)
		_shiftUp(heap, father_idx)
	}
}

// 比较复杂，还要考虑有没有左右子树
func _shiftDown(heap []int, idx int) {
	n := len(heap)
	lchild := idx * 2
	if lchild > n-1 {
		// 没有左子树 直接返回
		return
	}

	var min_child int
	rchild := lchild + 1
	if rchild > n-1 {
		// 没有右子树，
		min_child = lchild
	} else if heap[lchild] < heap[rchild] {
		min_child = lchild
	} else {
		min_child = rchild
	}
	if heap[idx] <= heap[min_child] {
		return
	} else {
		_swap(heap, idx, min_child)
		_shiftDown(heap, min_child)
	}
}
