package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	backPrint(arr, 0)
}

func backPrint(arr []int, cur int) {
	n := len(arr)
	left := n - cur
	if left == 0 {
		return
	}
	next := cur + 1
	backPrint(arr, next)
	fmt.Println(arr[cur])
}
