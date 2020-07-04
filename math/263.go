package main

import (
	"fmt"
)

// 263. 丑数
// 编写一个程序判断给定的数是否为丑数。
// 丑数就是只包含质因数 2, 3, 5 的正整数。
// 说明：
//     1 是丑数。
//     输入不会超过 32 位有符号整数的范围: [−231,  231 − 1]。
// https://leetcode-cn.com/problems/ugly-number/
func main() {
	fmt.Println(isUgly(6))
}

// 使用递归或迭代均可
func isUgly(num int) bool {
	if num > 0 {
		for i := 2; i < 6; i++ {
			for num%i == 0 {
				num /= i
			}
		}
	}
	return num == 1
}
