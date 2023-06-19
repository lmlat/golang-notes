package lc

import (
	"bytes"
	"fmt"
	"runtime"
	"unsafe"
)

/**
 *
 * @Author AiTao
 * @Date 2023/6/19 22:05
 * @Url
 **/

type Stack struct {
	bottom uintptr
	top    uintptr
	data   []unsafe.Pointer
}

func NewStack(values ...any) *Stack {
	s := &Stack{bottom: uintptr(0), top: uintptr(0)}
	if values == nil || len(values) == 0 {
		return s
	}
	for _, value := range values {
		s.Push(value)
	}
	return s
}

func (s *Stack) Push(value any) {
	if value == nil {
		return
	}
	ptr := unsafe.Pointer(&value)
	s.data = append(s.data, ptr)
	// 假设已经push了3个元素, 但又pop了3次, pop操作实际上只是将data中的指定元素设置为nil, 因此, 如果调用len(s.data)计算它的长度的话, 返回值仍然是3
	// 追加元素
	if s.top == uintptr(len(s.data)) {
		s.data = append(s.data, ptr)
	} else { // 重新设置值
		s.data[s.top] = ptr
	}
	s.top++
}

func (s *Stack) Pop() *any {
	if s.IsEmpty() {
		panic("empty")
	}
	val := s.data[s.top-1]
	s.top--
	s.data[s.top] = nil
	// 为了避免内存泄漏，需要在合适的时机释放栈中元素的内存。
	// 在出栈操作结束后，调用runtime.KeepAlive来保证被释放的元素仍然可以被访问一定的时间
	// KeepAlive函数会将其val标记为当前可访问。可确保在程序中调用 KeepAlive 的点之前不会释放对象，并且不会运行其终结器。
	runtime.KeepAlive(val)
	return toAnyPtr(val)
}
func (s *Stack) Peek() *any {
	if s.IsEmpty() {
		panic("empty")
	}
	return toAnyPtr(s.data[s.top-1])
}

func (s *Stack) PeekFirst() *any {
	if s.IsEmpty() {
		panic("empty")
	}
	return toAnyPtr(s.data[0])
}

// IsEmpty 判断栈是否为空
func (s *Stack) IsEmpty() bool {
	return s.top == s.bottom
}

// Size 获取栈大小
func (s *Stack) Size() int {
	return int(s.top - s.bottom)
}

// ForEach 遍历栈
func (s *Stack) ForEach(fn func(int, *any)) {
	if s.IsEmpty() || fn == nil {
		return
	}

	index := 0
	i := s.top - 1
	for i >= s.bottom && !s.IsEmpty() {
		fmt.Println("data:", s.data[i])
		fn(index, toAnyPtr(s.data[i]))
		if i == s.bottom {
			break
		}
		i--
		index++
	}
}

func (s *Stack) String() string {
	if s.IsEmpty() {
		return "[]"
	}
	buf := bytes.Buffer{}
	buf.WriteByte('[')
	s.ForEach(func(i int, val *any) {
		buf.WriteString(fmt.Sprintf("%v", *val))
		if (i + 1) != s.Size() {
			buf.WriteByte(',')
		}
	})
	buf.WriteByte(']')
	return buf.String()
}

func toAnyPtr(ptr unsafe.Pointer) *any {
	return (*any)(ptr)
}
