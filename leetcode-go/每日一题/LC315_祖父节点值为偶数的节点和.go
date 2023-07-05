package main

import (
	. "leetcode-go/utils"
)

/**
 *
 * @Author AiTao
 * @Date 2023/6/21 21:22
 * @Url https://leetcode.cn/problems/sum-of-nodes-with-even-valued-grandparent/
 **/

func main() {
}

// BFS
func sumEvenGrandparent(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	sum := 0
	for queue != nil {
		tmpQueue := queue
		queue = nil
		for _, node := range tmpQueue {
			if node.Left != nil {
				if node.Val%2 == 0 {
					if node.Left.Left != nil {
						sum += node.Left.Left.Val
					}
					if node.Left.Right != nil {
						sum += node.Left.Right.Val
					}
				}
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				if node.Val%2 == 0 {
					if node.Right.Left != nil {
						sum += node.Right.Left.Val
					}
					if node.Right.Right != nil {
						sum += node.Right.Right.Val
					}
				}
				queue = append(queue, node.Right)
			}
		}
	}
	return sum
}
