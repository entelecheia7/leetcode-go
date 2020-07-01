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

// 法一：暴力匹配。一旦匹配到一对括号，就进行消除。最后看这个字符串是否为空。O(n²)

// 法二：使用栈，遇到左括号，将右括号入栈，遇到右括号弹出匹配。O(n)
func isValid(s string) bool {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, ')')
		} else if s[i] == '[' {
			stack = append(stack, ']')
		} else if s[i] == '{' {
			stack = append(stack, '}')
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != s[i] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
