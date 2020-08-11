package main

import (
	"fmt"
)

// 696. 计数二进制子串
// 给定一个字符串 s，计算具有相同数量0和1的非空(连续)子字符串的数量，并且这些子字符串中的所有0和所有1都是组合在一起的。
// 重复出现的子串要计算它们出现的次数。
// 注意：
//     s.length 在1到50,000之间。
//     s 只包含“0”或“1”字符。
// https://leetcode-cn.com/problems/count-binary-substrings/
func main() {
	fmt.Println(countBinarySubstrings3("10101"))    // 4
	fmt.Println(countBinarySubstrings3("00110011")) // 6
}

// 法一：暴力
// 找到长度为2的合法子串，双指针向两侧扩展，计算数量
func countBinarySubstrings(s string) (count int) {
	n := len(s)
	if n <= 1 {
		return 0
	}

	for i := 1; i < n; i++ {
		if (s[i] == '0' && s[i-1] == '1') || (s[i] == '1' && s[i-1] == '0') {
			count++
			left, right := i-2, i+1
			for left >= 0 && right < n && s[left] == s[i-1] && s[right] == s[i] {
				count++
				left--
				right++
			}
		}
	}
	return count
}

// 法二：对连续相同字符进行分组
// 比如 00111011，分为[2 3 1 2]
// 遍历相邻的数对，计算总和
func countBinarySubstrings2(s string) (count int) {
	n := len(s)
	if n <= 1 {
		return 0
	}
	var group []int
	curGroup := 1
	for i := 1; i < n; i++ {
		if s[i] == s[i-1] {
			curGroup++
		} else {
			group = append(group, curGroup)
			curGroup = 1
		}
	}
	group = append(group, curGroup)
	n = len(group)
	for i := 1; i < n; i++ {
		count += getMin(group[i], group[i-1])
	}

	return
}

// 法三：对法二的空间优化
func countBinarySubstrings3(s string) (count int) {
	n := len(s)
	if n <= 1 {
		return 0
	}
	preGroup, curGroup := 0, 1
	for i := 1; i < n; i++ {
		if s[i] == s[i-1] {
			curGroup++
		} else {
			if preGroup != 0 {
				count += getMin((curGroup), preGroup)
			}
			preGroup = curGroup
			curGroup = 1
		}
	}
	count += getMin(preGroup, curGroup)
	return
}
func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
