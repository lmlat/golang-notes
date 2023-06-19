package lc

import (
	"bytes"
	"fmt"
)

/**
 *
 * @Author AiTao
 * @Date 2023/6/19 14:33
 * @Url
 **/

type Stack struct {
	data []any
}

// NewStack 初始化栈
// @param typ 栈中存储值的实际类型
func NewStack() *Stack {
	return &Stack{}
}

// Push 压栈
func (s *Stack) Push(value any) {
	if value == nil {
		return
	}
	s.data = append(s.data, value)
}

// Pop 出栈
// @return 返回栈顶元素, 并删除
func (s *Stack) Pop() any {
	if s.IsEmpty() {
		return nil
	}
	idx := len(s.data) - 1
	oldVal := s.data[idx]
	s.data = s.data[:idx]
	return oldVal
}

// PeekLast 获取栈顶元素
// @return 返回栈顶元素
func (s *Stack) PeekLast() any {
	if s.IsEmpty() {
		return nil
	}
	return s.data[s.Size()-1]
}

func (s *Stack) PeekFirst() any {
	if s.IsEmpty() {
		return nil
	}
	return s.data[0]
}

// IsEmpty 判断栈是否为空
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

// Size 获取栈大小
func (s *Stack) Size() int {
	return len(s.data)
}

// ForEach 遍历栈
func (s *Stack) ForEach(fn func(int, any)) {
	if fn == nil {
		return
	}

	for i, val := range s.data {
		fn(i, val)
	}
}

func (s *Stack) String() string {
	buf := bytes.Buffer{}
	buf.WriteByte('[')
	s.ForEach(func(i int, val any) {
		buf.WriteString(fmt.Sprintf("%v", val))
		if (i + 1) != s.Size() {
			buf.WriteByte(',')
		}
	})
	buf.WriteByte(']')
	return buf.String()
}
