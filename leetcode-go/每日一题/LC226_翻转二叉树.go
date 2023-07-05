package main

import . "leetcode-go/utils"

/**
 *
 * @Author AiTao
 * @Date 2023/6/22 0:20
 * @Url
 **/

// 前序遍历, 直接交换节点
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	tmp := root.Right
	root.Right = root.Left
	root.Left = tmp
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

func main() {
}
