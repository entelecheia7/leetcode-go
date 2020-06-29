package main

import (
	"fmt"
)

// 20. 有效的括号
// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
// 有效字符串需满足：
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 注意空字符串可被认为是有效字符串。
// https://leetcode-cn.com/problems/valid-parentheses
func main() {
	// fmt.Println(isValid("{]"))
	fmt.Println(isValid("()"))
}

// 使用栈，遇到左括号入栈，遇到右括号弹出匹配
func isValid(s string) bool {
	m := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	n := len(s)
	if n == 0 {
		return true
	} else if n%2 == 1 {
		return false
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		switch s[i] {
		case '(', '[', '{':
			stack = append(stack, s[i])
		default:
			if len(stack) == 0 || stack[len(stack)-1] != m[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
