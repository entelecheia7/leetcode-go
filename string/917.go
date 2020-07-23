package main

import "fmt"

// 917. 仅仅反转字母
// 给定一个字符串 S，返回 “反转后的” 字符串，其中不是字母的字符都保留在原地，而所有字母的位置发生反转。
// 提示：
//     S.length <= 100
//     33 <= S[i].ASCIIcode <= 122
// 	S 中不包含 \ or "
// https://leetcode-cn.com/problems/reverse-only-letters/
func main() {
	fmt.Println(reverseOnlyLetters("a-bC-dEf-ghIj")) // "j-Ih-gfE-dCba"
}

// 双指针，O(n)
func reverseOnlyLetters(S string) string {
	length := len(S)
	if length <= 1 {
		return S
	}
	bytes := []byte(S)
	left, right := 0, length-1
	for left < right {
		if !isLetter(bytes[left]) {
			left++
		} else if !isLetter(bytes[right]) {
			right--
		} else {
			bytes[left], bytes[right] = bytes[right], bytes[left]
			left++
			right--
		}
	}
	return string(bytes)
}

func isLetter(x byte) bool {
	if x >= 'a' && x <= 'z' {
		return true
	} else if x >= 'A' && x <= 'Z' {
		return true
	}

	return false
}
