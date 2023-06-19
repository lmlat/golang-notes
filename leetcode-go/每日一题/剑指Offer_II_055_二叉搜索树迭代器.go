package main

import (
	"fmt"
	. "leetcode-go/utils"
)

/**
 *
 * @Author AiTao
 * @Date 2023/6/17 21:42
 * @Url https://leetcode.cn/problems/kTOapQ/
 **/

type BSTIterator struct {
	cursor int
	values []int
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{cursor: 0, values: infixOrder(root)}
}

func infixOrder(root *TreeNode) []int {
	values := make([]int, 0)
	if root == nil {
		return values
	}
	values = append(values, infixOrder(root.Left)...)
	values = append(values, root.Val)
	values = append(values, infixOrder(root.Right)...)
	return values
}

func (this *BSTIterator) Next() int {
	value := this.values[this.cursor]
	this.cursor++
	return value
}

func (this *BSTIterator) HasNext() bool {
	return this.cursor < len(this.values)
}

type BSTIterator2 struct {
	stack []*TreeNode
	cur   *TreeNode
}

func Constructor2(root *TreeNode) BSTIterator2 {
	return BSTIterator2{cur: root}
}

func (this *BSTIterator2) Next() int {
	// 将所有左节点压栈
	for node := this.cur; node != nil; node = node.Left {
		this.stack = append(this.stack, node)
	}
	size := len(this.stack) - 1
	this.cur = this.stack[size]
	// 出栈
	this.stack = this.stack[:size]
	curVal := this.cur.Val
	/*
		如果当前遍历的左节点中不存在左节点, 但存在右节点
		      7
		    /   \
		   3     15
		    \    / \
		     4  9   20
	*/
	this.cur = this.cur.Right
	return curVal
}

func (this *BSTIterator2) HasNext() bool {
	return this.cur != nil || len(this.stack) > 0
}

func main() {
	root := &TreeNode{Val: 7,
		Left:  &TreeNode{3, nil, nil},
		Right: &TreeNode{15, &TreeNode{9, nil, nil}, &TreeNode{20, nil, nil}}}
	iterator := Constructor2(root)
	fmt.Println(iterator.Next())    // 3
	fmt.Println(iterator.Next())    // 7
	fmt.Println(iterator.HasNext()) // true
	fmt.Println(iterator.Next())    // 9
	fmt.Println(iterator.HasNext()) // true
	fmt.Println(iterator.Next())    // 15
	fmt.Println(iterator.HasNext()) // true
	fmt.Println(iterator.Next())    // 20
	fmt.Println(iterator.HasNext()) // true
}
