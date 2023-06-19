package main

import (
	"fmt"
	. "leetcode-go/utils"
)

/**
 *
 * @Author AiTao
 * @Date 2023/6/18 17:38
 * @Url https://leetcode.cn/problems/count-nodes-equal-to-average-of-subtree/
 **/
func main() {
	root := &TreeNode{4,
		&TreeNode{8,
			&TreeNode{0, nil, nil},
			&TreeNode{1, nil, nil}},
		&TreeNode{5, nil, &TreeNode{6, nil, nil}},
	}
	fmt.Println(averageOfSubtree(root))
	fmt.Println(averageOfSubtree2(root))
}

// dfs
func averageOfSubtree(root *TreeNode) (ans int) {
	if root == nil {
		return 0
	}

	var dfs func(*TreeNode) (int, int)
	dfs = func(root *TreeNode) (sum int, count int) {
		if root == nil {
			return
		}
		// 计算左子树的总和和节点数
		leftSum, leftCount := dfs(root.Left)
		// 计算右子树的总和和节点数
		rightSum, rightCount := dfs(root.Right)
		// 求和
		sum = leftSum + rightSum + root.Val
		// 求总数
		count = leftCount + rightCount + 1
		// 符合条件
		if root.Val == sum/count {
			ans++
		}
		return
	}
	dfs(root)
	return
}

// 栈, 只能过93/138个用例
func averageOfSubtree2(root *TreeNode) (ans int) {
	if root == nil {
		return 0
	}

	stack := NewStack()
	stack.Push(root)
	callSumAndCount := func(node *TreeNode) (int, int) {
		stack.Push(node)
		curSum, curCount := node.Val, 1
		if node.Left != nil {
			curSum += node.Left.Val
			curCount++
		}
		if node.Right != nil {
			curSum += node.Right.Val
			curCount++
		}
		return curSum, curCount
	}

	for !stack.IsEmpty() {
		// 出栈
		node := stack.Pop().(*TreeNode)
		// 计算当前节点的总和和节点数
		sum, count := node.Val, 1
		if node.Left != nil {
			leftSum, leftCount := callSumAndCount(node.Left)
			sum += leftSum
			count += leftCount
		}
		if node.Right != nil {
			rightSum, rightCount := callSumAndCount(node.Right)
			sum += rightSum
			count += rightCount
		}
		// 判断当前节点是否符合条件
		if node.Val == sum/count {
			ans++
		}
	}
	return
}
