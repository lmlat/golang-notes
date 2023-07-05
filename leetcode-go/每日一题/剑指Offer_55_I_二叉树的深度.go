package main

import (
	. "leetcode-go/utils"
	"math"
)

/**
 *
 * @Author AiTao
 * @Date 2023/6/22 0:15
 * @Url https://leetcode.cn/problems/er-cha-shu-de-shen-du-lcof/
 **/

func main() {
}

func maxDepth0(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return int(math.Max(float64(maxDepth0(root.Left)), float64(maxDepth0(root.Right))) + 1)
}
