package main

import (
	"algorithm/sort"
	"algorithm/utils"
	"time"
)

func maxSubBackStr(s string) string {
	n := len(s)
	if n == 1 {
		return s
	}
	maxstr := ""
	for i := 0; i <= n-2; i++ {
		str1 := findByMiddle(s, i, i)
		if len(str1) > len(maxstr) {
			maxstr = str1
		}
		str2 := findByMiddle(s, i, i+1)
		if len(str2) > len(maxstr) {
			maxstr = str2
		}
	}
	return maxstr
}

// 找出 以 [l,r] 为中心的最长回文串
func findByMiddle(str string, l, r int) string {
	if str[l] != str[r] {
		return string(str[l])
	}
	var (
		n = len(str)
	)
	for l >= 0 && r <= n-1 {
		if str[l] == str[r] {
			l--
			r++
		} else {
			break
		}
	}
	return str[l+1 : r]
}

func findMin(arr []int) (int, int) {
	var (
		min_val = arr[0]
		min_idx = 0
	)
	for idx, val := range arr {
		if val < min_val {
			min_idx = idx
			min_val = val
		}
	}
	return min_idx, min_val
}
func getFinalState(nums []int, k int, multiplier int) []int {
	if k <= 0 {
		return nums
	}

	utils.VarDump(nums)
	for i := 1; i <= k; i++ {
		utils.VarDump(i)
		min_idx, min_val := findMin(nums)
		nums[min_idx] = min_val * multiplier
		utils.VarDump(nums)
	}
	utils.VarDump(nums)
	return nums
}

func findMid(a, b []int) float64 {
	utils.VarDump(a, b)
	n := len(a)
	m := len(b)

	sum := n + m

	if sum == 1 {
		if n == 0 {
			return float64(b[0])
		} else {
			return float64(a[0])
		}
	} else if sum == 2 {
		if n == 0 {
			return (float64(b[0]) + float64(b[1])) / 2
		} else if n == 1 {
			return (float64(a[0]) + float64(b[0])) / 2
		} else {
			return (float64(a[0]) + float64(a[1])) / 2
		}
	} else if n == 0 {
		// 超过2个 还需要减少
		// 找到min和max
		// 有可能全部在 a上
		return findMid(a, b[1:m-1])
	} else if m == 0 {
		// 或者b上
		return findMid(a[1:n-1], b)
	} else {
		// a 和 b 上都有，并且
		// 去掉最小
		if a[0] < b[0] {
			a = a[1:]
			// 去掉最大
			if len(a) == 0 {
				b = b[:m-1]
			} else {
				if a[len(a)-1] > b[len(b)-1] {
					a = a[:len(a)-1]
				} else {
					b = b[:len(b)-1]
				}
			}
		} else {
			b = b[1:]
			// 去掉最大
			if len(b) == 0 {
				a = a[:n-1]
			} else {
				if a[len(a)-1] > b[len(b)-1] {
					a = a[:len(a)-1]
				} else {
					b = b[:len(b)-1]
				}
			}
		}
		return findMid(a, b)
	}
}

func in(str string, mmap map[string]int) bool {
	num, exist := mmap[str]
	return exist && num > 0
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	i := 0
	j := 0
	charMap := map[string]int{}
	mmax := 0 // 最大值
	StrArr := []string{}
	for j <= n-1 {
		curStr := string(s[j])
		if !in(curStr, charMap) {
			// 执行+的操作
			StrArr = append(StrArr, curStr)
			charMap[curStr] = 1 // 当前字符写入

			mmax = max(mmax, len(StrArr))
			j++
		} else {
			// -i 的操作
			position_i_str := StrArr[0]
			StrArr = StrArr[1:] // 去掉第一个元素
			charMap[position_i_str]--
			i++
		}
		utils.VarDump(i, j, mmax, charMap, StrArr)

		time.Sleep(time.Second)
	}
	return mmax
}

func timeTest() {
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
	arr := utils.RandIntArr(0, 10, 10)
	utils.VarDump(arr)
	f(arr)
	utils.VarDump(arr)
	if err := utils.Check(arr); err != nil {
		utils.VarDump(err)
	}
}
