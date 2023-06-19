package main

import (
	"fmt"
	. "leetcode-go/utils"
)

/**
 *
 * @Author AiTao
 * @Date 2023/6/10 23:46
 * @Url https://leetcode.cn/problems/count-complete-tree-nodes/
 **/

func calcSize(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return calcSize(root.Left) + calcSize(root.Right) + 1
}
func countNodes(root *TreeNode) int {
	return calcSize(root)
}

// 利用满二叉树性质的解法
func countNodes2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := calcHeight(root.Left)
	rightHeight := calcHeight(root.Right)
	if leftHeight == rightHeight {
		return countNodes2(root.Right) + (1 << leftHeight)
	} else {
		return countNodes2(root.Left) + (1 << rightHeight)
	}
}

func calcHeight(node *TreeNode) int {
	height := 0
	for node != nil {
		height++
		node = node.Left
	}
	return height
}

func main() {
	root := &TreeNode{Val: 4,
		Left: &TreeNode{Val: 1, Left: &TreeNode{},
			Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
		}, Right: &TreeNode{Val: 6, Left: &TreeNode{Val: 5},
			Right: &TreeNode{Val: 7, Right: &TreeNode{Val: 8}},
		}}
	fmt.Println(countNodes(root))
	fmt.Println(countNodes2(root))
}
