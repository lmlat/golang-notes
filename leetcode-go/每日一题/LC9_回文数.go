package main

import "fmt"

/**
 *
 * @Author AiTao
 * @Date 2023/6/9 22:05
 * @Url https://leetcode.cn/problems/palindrome-number/
 **/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	sum := 0
	y := x
	for y != 0 {
		sum = sum*10 + (y % 10)
		y /= 10
	}

	return sum == x
}
func main() {
	fmt.Println(isPalindrome(121))
}
