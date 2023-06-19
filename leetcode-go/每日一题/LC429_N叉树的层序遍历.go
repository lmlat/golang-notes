package main

import (
	. "leetcode-go/utils"
)

/**
 *
 * @Author AiTao
 * @Date 2023/6/13 23:25
 * @Url
 **/
func levelOrder(root *Node) (ret [][]int) {
	if root == nil {
		return
	}
	queue := []*Node{root}
	for queue != nil {
		var level []int
		tmpQueue := queue
		queue = nil
		for _, node := range tmpQueue {
			level = append(level, node.Val)
			queue = append(queue, node.Children...)
		}
		ret = append(ret, level)
	}
	return
}

func main() {
}
