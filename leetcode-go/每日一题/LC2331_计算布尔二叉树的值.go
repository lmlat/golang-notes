package main

import . "leetcode-go/utils"

/**
 *
 * @Author AiTao
 * @Date 2023/6/18 0:09
 * @Url https://leetcode.cn/problems/evaluate-boolean-binary-tree/
 **/

// 前序遍历
func evaluateTree(root *TreeNode) bool {
	if root.Left == nil {
		return root.Val == 1
	}
	// 逻辑或
	if root.Val == 2 {
		return evaluateTree(root.Left) || evaluateTree(root.Right)
	}
	// 逻辑与
	return evaluateTree(root.Left) && evaluateTree(root.Right)
}

func main() {
}
