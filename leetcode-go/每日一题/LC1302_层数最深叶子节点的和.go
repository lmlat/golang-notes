package main

import (
	"fmt"
	. "leetcode-go/utils"
)

/**
 *
 * @Author AiTao
 * @Date 2023/6/17 20:49
 * @Url
 **/

func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var sum int
	queue := []*TreeNode{root}
	for queue != nil {
		// 计算每层节点的和
		sum = 0
		tmpQueue := queue
		queue = nil
		for _, node := range tmpQueue {
			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return sum
}

func main() {
	root := &TreeNode{1,
		&TreeNode{2,
			&TreeNode{4, &TreeNode{7, nil, nil},
				nil}, &TreeNode{5, nil, nil}},
		&TreeNode{3, nil, &TreeNode{6, nil, &TreeNode{8, nil, nil}}}}
	fmt.Println(deepestLeavesSum(root))
}
