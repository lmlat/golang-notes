package main

import (
	. "leetcode-go/utils"
)

/**
 *
 * @Author AiTao
 * @Date 2023/6/14 22:18
 * @Url https://leetcode.cn/problems/KnLfVT/
 **/

// BFS
func expandBinaryTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	queue := []*TreeNode{root}
	for queue != nil {
		tmpQueue := queue
		queue = nil
		for _, node := range tmpQueue {
			if node.Left != nil {
				queue = append(queue, node.Left)
				node.Left = &TreeNode{Val: -1, Left: node.Left}
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
				node.Right = &TreeNode{Val: -1, Right: node.Right}
			}
		}
	}

	return root
}

func main() {
	root := &TreeNode{Val: 3,
		Left:  &TreeNode{1, &TreeNode{3, nil, nil}, &TreeNode{8, nil, nil}},
		Right: &TreeNode{7, nil, &TreeNode{4, nil, nil}}}

	expandBinaryTree(root)
}
