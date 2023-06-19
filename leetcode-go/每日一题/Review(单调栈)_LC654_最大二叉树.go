package main

import (
	"fmt"
	. "leetcode-go/utils"
)

/**
 *
 * @Author AiTao
 * @Date 2023/6/19 19:51
 * @Url https://leetcode.cn/problems/maximum-binary-tree/
 */
func main() {
	nums := []int{3, 2, 1, 6, 0, 5}
	forEach(constructMaximumBinaryTree2(nums))
}

func forEach(root *TreeNode) {
	if root != nil {
		fmt.Println(root.Val)
		forEach(root.Left)
		forEach(root.Right)
	}
}

// dfs
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	val, i := getMaxValueAndIndex(nums)
	return &TreeNode{
		val,
		constructMaximumBinaryTree(nums[:i]),
		constructMaximumBinaryTree(nums[i+1:]),
	}
}

// 获取切片中的最大值和索引值
func getMaxValueAndIndex(nums []int) (maxVal int, index int) {
	maxVal, index = nums[0], 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxVal {
			maxVal, index = nums[i], i
		}
	}
	return
}

// 单调栈(递增)
func constructMaximumBinaryTree2(nums []int) *TreeNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	var stack []*TreeNode
	for i := 0; i < len(nums); i++ {
		node := &TreeNode{Val: nums[i]}
		for len(stack) > 0 {
			// 获取栈顶元素
			top := stack[len(stack)-1]
			// 如果栈顶元素大于待插入元素, 直接入栈
			if top.Val > node.Val {
				stack = append(stack, node)
				top.Right = node
				break
			}

			// 出栈
			stack = stack[:len(stack)-1]
			node.Left = top
		}
		if len(stack) == 0 {
			stack = append(stack, node)
		}
	}
	return stack[0]
}
