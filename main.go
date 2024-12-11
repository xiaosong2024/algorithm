package main

import (
	"algorithm/sort"
	"algorithm/utils"
	"time"
)

func main() {
	go func() {
		t := time.Now()
		for i := 0; i < 1000; i++ {
			test(sort.BubbleSort)
		}
		utils.VarDump("BubbleSort", time.Now().Sub(t).Seconds())
	}()

	go func() {
		t := time.Now()
		for i := 0; i < 1000; i++ {
			test(sort.SelectSort)
		}
		utils.VarDump("SelectSort", time.Now().Sub(t).Seconds())
	}()

	go func() {
		t := time.Now()
		for i := 0; i < 1000; i++ {
			test(sort.InsertSort)
		}
		utils.VarDump("InsertSort", time.Now().Sub(t).Seconds())
	}()

	go func() {
		t := time.Now()
		for i := 0; i < 1000; i++ {
			test(sort.MergeSort)
		}
		utils.VarDump("MergeSort", time.Now().Sub(t).Seconds())
	}()

	go func() {
		t := time.Now()
		for i := 0; i < 1000; i++ {
			test(sort.QuickSort)
		}
		utils.VarDump("QuickSort", time.Now().Sub(t).Seconds())
	}()
	for {
	}

}

func test(f func([]int)) {
	arr := utils.RandIntArr(0, 10000, 10000)
	f(arr)
	if err := utils.Check(arr); err != nil {
		utils.VarDump(err)
	}
}
