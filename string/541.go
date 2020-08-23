package main

import "fmt"

// 541. 反转字符串 II
// 给定一个字符串 s 和一个整数 k，你需要对从字符串开头算起的每隔 2k 个字符的前 k 个字符进行反转。
//     如果剩余字符少于 k 个，则将剩余字符全部反转。
//     如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。
// 	提示：
//     该字符串只包含小写英文字母。
//     给定字符串的长度和 k 在 [1, 10000] 范围内。
// https://leetcode-cn.com/problems/reverse-string-ii/
func main() {
	fmt.Println(reverseStr("abcdefg", 2)) // bacdfeg
}

// 法一：暴力
func reverseStr(s string, k int) string {
	n := len(s)
	result := make([]byte, 0, n)
	start := 0
	reverseFlg := true
	for start < n {
		right := getMin(n-1, start+k-1)
		if reverseFlg {
			for i := right; i >= start; i-- {
				result = append(result, s[i])
			}
		} else {
			result = append(result, []byte(s[start:right+1])...)
		}
		reverseFlg = !reverseFlg
		start += k
	}

	return string(result)
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
