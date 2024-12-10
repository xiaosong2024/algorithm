package main

import (
	"algorithm/sort"
	"algorithm/tree"
	"algorithm/utils"
	"fmt"
)

func main() {

	tree1 := &tree.BiTree{
		Value: 1,
		LeftChild: &tree.BiTree{
			Value: 2,
			LeftChild: &tree.BiTree{
				Value:      4,
				LeftChild:  nil,
				RightChild: nil,
			},
			RightChild: &tree.BiTree{
				Value:      5,
				LeftChild:  nil,
				RightChild: nil,
			},
		},
		RightChild: &tree.BiTree{
			Value: 3,
			LeftChild: &tree.BiTree{
				Value:      6,
				LeftChild:  nil,
				RightChild: nil,
			},
			RightChild: &tree.BiTree{
				Value:      7,
				LeftChild:  nil,
				RightChild: nil,
			},
		},
	}
	tree.Floor(tree1)
}

func test() {
	arr := utils.RandIntArr(0, 100, 10)
	sort.InsertSort(arr)
	utils.VarDump(arr)
	fmt.Println(utils.Check(arr))
}
