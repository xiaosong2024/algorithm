package main

import (
	"algorithm/sort"
	"algorithm/utils"
	"fmt"
)

func main() {
	for i := 0; i <= 10; i++ {
		test()
	}
}

func test() {
	arr := utils.RandIntArr(0, 100, 10)
	sort.SelectSort(arr)
	utils.VarDump(arr)
	fmt.Println(utils.Check(arr))
}
