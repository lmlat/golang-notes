package main

import . "leetcode-go/utils"

/**
 *
 * @Author AiTao
 * @Date 2023/6/22 0:37
 * @Url
 **/

func main() {
}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	tmp := root.Right
	root.Right = root.Left
	root.Left = tmp
	mirrorTree(root.Left)
	mirrorTree(root.Right)
	return root
}