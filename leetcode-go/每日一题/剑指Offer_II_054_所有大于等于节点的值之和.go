package main

import (
	"fmt"
	. "leetcode-go/utils"
)

/**
 *
 * @Author AiTao
 * @Date 2023/6/9 22:25
 * @Url https://leetcode.cn/problems/w6cpku/
 **/

func main() {
	root := &TreeNode{Val: 4,
		Left: &TreeNode{Val: 1, Left: &TreeNode{},
			Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
		}, Right: &TreeNode{Val: 6, Left: &TreeNode{Val: 5},
			Right: &TreeNode{Val: 7, Right: &TreeNode{Val: 8}},
		}}
	convertBST(root)
	foreach(root, func(node *TreeNode) {
		fmt.Println(node.Val)
	})
}

// 暴力
func convertBST(root *TreeNode) *TreeNode {
	m := make(map[int]int)
	foreach(root, func(n1 *TreeNode) {
		replaceVal := n1.Val
		foreach(root, func(n2 *TreeNode) {
			if n2.Val > n1.Val {
				replaceVal += n2.Val
			}
		})
		m[n1.Val] = replaceVal
	})
	foreach(root, func(n1 *TreeNode) {
		n1.Val = m[n1.Val]
	})
	return root
}
func foreach(node *TreeNode, fn func(*TreeNode)) {
	if node != nil {
		fn(node)
		foreach(node.Left, fn)
		foreach(node.Right, fn)
	}
}

// 逆中序遍历
func convertBST2(root *TreeNode) *TreeNode {
	sum := 0
	var foreach2 func(*TreeNode)
	foreach2 = func(node *TreeNode) {
		if node != nil {
			// 先遍历右子树
			foreach2(node.Right)
			sum += node.Val
			node.Val = sum
			// 再遍历左子树
			foreach2(node.Left)
		}
	}
	foreach2(root)
	return root
}

// Morris[ˈmɔrɪs]遍历
func convertBST3(root *TreeNode) *TreeNode {

	return nil
}
