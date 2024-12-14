package tree

import "algorithm/utils"

type BiTree struct {
	Value      any
	LeftChild  *BiTree
	RightChild *BiTree
}

func Front(tree *BiTree) {
	if tree == nil {
		return
	}
	utils.VarDump(tree.Value)
	Front(tree.LeftChild)
	Front(tree.RightChild)
}

func Middle(tree *BiTree) {
	if tree == nil {
		return
	}
	Middle(tree.LeftChild)
	utils.VarDump(tree.Value)
	Middle(tree.RightChild)
}

func Back(tree *BiTree) {
	if tree == nil {
		return
	}
	Back(tree.LeftChild)
	Back(tree.RightChild)
	utils.VarDump(tree.Value)
}

// Floor 层序遍历
// 读取自己，将左右子树入队，然后从队列取出一个，继续
