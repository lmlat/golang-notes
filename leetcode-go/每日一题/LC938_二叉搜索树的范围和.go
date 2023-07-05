package main

import . "leetcode-go/utils"

/**
 *
 * @Author AiTao
 * @Date 2023/6/21 20:49
 * @Url https://leetcode.cn/problems/range-sum-of-bst/
 **/

func main() {
}

// 中序遍历
func rangeSumBST(root *TreeNode, low int, high int) int {
	sum := 0
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		if root.Val >= low && root.Val <= high {
			sum += root.Val
		}
		dfs(root.Right)
	}
	dfs(root)
	return sum
}
