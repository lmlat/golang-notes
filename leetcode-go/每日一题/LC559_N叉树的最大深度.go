package main

import . "leetcode-go/utils"

/**
 *
 * @Author AiTao
 * @Date 2023/6/13 23:44
 * @Url https://leetcode.cn/problems/maximum-depth-of-n-ary-tree/
 **/

func maxDepth(root *Node) (height int) {
	if root == nil {
		return 0
	}
	queue := []*Node{root}
	for queue != nil {
		tmpQueue := queue
		queue = nil
		for _, node := range tmpQueue {
			queue = append(queue, node.Children...)
		}
		height++
	}
	return
}

func main() {
}
