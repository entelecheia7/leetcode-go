package main

import "fmt"

// 9. 回文数
// 判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
// 进阶:
// 你能不将整数转为字符串来解决这个问题吗？
// https://leetcode-cn.com/problems/palindrome-number/
func main() {
	fmt.Println(isPalindrome(1001))
}

// 法一：转换为字符串，双指针，略

// 法二：通过个位数去计算回文数，最后判断是否相等
func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) { // 初始x的末尾不能是0
		return false
	} else if x < 10 {
		return true
	}
	x2 := x
	y := 0
	for x2 > 0 {
		y = y*10 + x2%10
		x2 /= 10
	}
	return y == x
}

// 法二的另一种写法
// 但lc上的时间反而更长了
func isPalindrome2(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) { // 初始x的末尾不能是0
		return false
	} else if x < 10 {
		return true
	}
	y := 0
	for x > y {
		y = y*10 + x%10
		x /= 10
	}
	// 考虑x的位数是奇数、偶数两种情况
	return y == x || x == y/10
}
