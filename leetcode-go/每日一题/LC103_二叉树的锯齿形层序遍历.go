package main

import (
	"fmt"
	. "leetcode-go/utils"
)

/**
 *
 * @Author AiTao
 * @Date 2023/6/14 23:07
 * @Url
 **/

func zigzagLevelOrder(root *TreeNode) (ret [][]int) {
	if root == nil {
		return
	}
	// false:从左向右, true:从右向左
	flag := false
	queue := []*TreeNode{root}
	for queue != nil {
		tmpQueue := queue
		queue = nil
		var level []int
		fmt.Println("flag:", flag)
		for _, node := range tmpQueue {
			if !flag {
				level = append(level, node.Val)
			} else {
				level = append([]int{node.Val}, level...)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		flag = !flag
		ret = append(ret, level)
	}
	return
}

func appendElement(queue []*TreeNode, node *TreeNode, flag bool) (ret []*TreeNode) {
	if !flag {
		ret = append([]*TreeNode{node}, queue...)

	} else {
		ret = append(queue, node)
	}
	return
}

func main() {
	root := &TreeNode{Val: 3,
		Left:  &TreeNode{9, nil, nil},
		Right: &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}

	fmt.Println(zigzagLevelOrder(root))
}
