package main

import (
	"big"
	"fmt"
	"strconv"
)

// 67. 二进制求和
// 给你两个二进制字符串，返回它们的和（用二进制表示）。
// 输入为 非空 字符串且只包含数字 1 和 0。
// 提示：
// 每个字符串仅由字符 '0' 或 '1' 组成。
// 1 <= a.length, b.length <= 10^4
// 字符串如果不是 "0" ，就都不含前导零。
// https://leetcode-cn.com/problems/add-binary/
func main() {
	fmt.Println(addBinary("1010", "1011"))
	fmt.Println(addBinary("11", "1"))
}

// 法一：按照进制计算的定义运算（需要使用除法
func addBinary(a string, b string) string {
	la := len(a)
	lb := len(b)
	n := getMax(la, lb)
	s := ""
	carry := 0 // 进位标识

	for i := 0; i < n; i++ {
		if i < la {
			carry += int(a[la-i-1] - '0')
		}
		if i < lb {
			carry += int(b[lb-i-1] - '0')
		}
		s = strconv.Itoa(carry%2) + s
		carry /= 2
	}
	if carry > 0 {
		s = "1" + s
	}

	return s
}
func getMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// 法二：位运算
func addBinary2(a string, b string) string {
	a1, _ := new(big.Int).SetString(a, 2) // 保存进位
	b1, _ := new(big.Int).SetString(b, 2)
	var zero big.Int
	for a1.Cmp(&zero) > 0 {
		res := new(big.Int).Xor(a1, b1)
		carry := b1.And(a1, b1).Lsh(b1, 1)
		a1, b1 = carry, res
	}
	return fmt.Sprintf("%b", x)
}
