package main

import (
	"fmt"
)

// 367. 有效的完全平方数
// 给定一个正整数 num，编写一个函数，如果 num 是一个完全平方数，则返回 True，否则返回 False。
// 说明：不要使用任何内置的库函数，如  sqrt。
// https://leetcode-cn.com/problems/valid-perfect-square/
func main() {
	fmt.Println(isPerfectSquare(16))
	fmt.Println(isPerfectSquare2(16)) //best
}

// 可以转换为求平方根问题
// 法一：二分法
func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	left, right := 1, num
	for left <= right {
		mid := left + ((right - left) >> 1)
		if mid*mid <= num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return r*r == num
}

// 法二：牛顿迭代法
// 收敛的O(logn)，优于二分法，空间复杂度更低
// best
func isPerfectSquare2(num int) bool {
	if num == 1 {
		return true
	}
	y := num >> 1
	for y*y > num {
		y = (y + num/y) >> 1
	}
	return y*y == num
}

// 此外还有袖珍计算器算法和位运算，参考 69. x 的平方根
