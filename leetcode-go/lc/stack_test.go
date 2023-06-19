package lc_test

import (
	"fmt"
	. "leetcode-go/lc"
	"testing"
)

/**
 *
 * @Author AiTao
 * @Date 2023/6/19 23:37
 * @Url
 **/

func TestNewStack(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(*s.Peek())
	fmt.Println(s.Size())
	fmt.Println("IsEmpty:", s.IsEmpty())
	fmt.Println(s.String())
	fmt.Println("pop:", *s.Pop())
	fmt.Println(s.String())
	fmt.Println("pop:", *s.Pop())
	fmt.Println(s.String())
	fmt.Println("pop:", *s.Pop())
	fmt.Println("IsEmpty:", s.IsEmpty())
}
