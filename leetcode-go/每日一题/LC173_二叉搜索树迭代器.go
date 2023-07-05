package main

import . "leetcode-go/utils"

/**
 *
 * @Author AiTao
 * @Date 2023/6/21 20:59
 * @Url https://leetcode.cn/problems/binary-search-tree-iterator/
 **/

type BSTIterator3 struct {
	stack []*TreeNode
	cur   *TreeNode
}

func Constructor3(root *TreeNode) BSTIterator3 {
	return BSTIterator3{cur: root}
}

func (this *BSTIterator3) Next3() int {
	for node := this.cur; node != nil; node = node.Left {
		this.stack = append(this.stack, node)
	}
	size := len(this.stack)
	this.cur = this.stack[size-1]
	this.stack = this.stack[:size]
	curVal := this.cur.Val
	this.cur = this.cur.Right
	return curVal
}

func (this *BSTIterator3) HasNext3() bool {
	return this.cur != nil || len(this.stack) > 0
}

func main() {
}
