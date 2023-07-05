package main

import (
	"fmt"
	. "leetcode-go/utils"
)

/**
 *
 * @Author AiTao
 * @Date 2023/6/21 22:12
 * @Url
 **/

func main() {
	head := &ListNode{-10, &ListNode{-3, &ListNode{0, &ListNode{5, &ListNode{9, nil}}}}}
	For2(sortedListToBST(head))
}
func For2(root *TreeNode) {
	if root != nil {
		fmt.Println(root.Val)
		For2(root.Left)
		For2(root.Right)
	}
}

/*
思路：升序链表是中序遍历的结果, 根据中序遍历方式来创建节点
*/
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	// 计算链表的个数
	length := 0
	cur := head
	for cur != nil {
		length++
		cur = cur.Next
	}
	// 构建BST
	var buildBST func(int, int) *TreeNode
	buildBST = func(start, end int) *TreeNode {
		if start > end {
			return nil
		}
		mid := (start + end) >> 1
		root := &TreeNode{}
		// 左节点
		root.Left = buildBST(start, mid-1)
		// 中序遍历节点值直接当前节点值
		root.Val = head.Val
		// 中序遍历下一个值
		head = head.Next
		// 右节点
		root.Right = buildBST(mid+1, end)
		return root
	}
	return buildBST(0, length-1)
}
