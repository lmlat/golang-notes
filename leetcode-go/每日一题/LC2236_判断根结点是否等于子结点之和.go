package main

import . "leetcode-go/utils"

/**
 *
 * @Author AiTao
 * @Date 2023/6/17 20:41
 * @Url https://leetcode.cn/problems/root-equals-sum-of-children/
 **/
func checkTree(root *TreeNode) bool {
	return root.Val == root.Left.Val+root.Right.Val
}
