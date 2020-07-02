package main

import (
	"fmt"
)

// 125. 删除最外层的括号
// 有效括号字符串为空 ("")、"(" + A + ")" 或 A + B，其中 A 和 B 都是有效的括号字符串，+ 代表字符串的连接。例如，""，"()"，"(())()" 和 "(()(()))" 都是有效的括号字符串。
// 如果有效字符串 S 非空，且不存在将其拆分为 S = A+B 的方法，我们称其为原语（primitive），其中 A 和 B 都是非空有效括号字符串。
// 给出一个非空有效字符串 S，考虑将其进行原语化分解，使得：S = P_1 + P_2 + ... + P_k，其中 P_i 是有效括号字符串原语。
// 对 S 进行原语化分解，删除分解中每个原语字符串的最外层括号，返回 S 。
// 提示：
//     S.length <= 10000
//     S[i] 为 "(" 或 ")"
//     S 是一个有效括号字符串
// 链接：https://leetcode-cn.com/problems/remove-outermost-parentheses
func main() {
	fmt.Println(removeOuterParentheses2("(()())(())")) //best
}

// 法一：使用栈，遇到'('入栈，')'出栈。当栈为空时，获得当前原语的位置，去掉外层括号
// O(n)
func removeOuterParentheses(S string) (result string) {
	if S == "" {
		return S
	}
	start := 0 // 上一个原语的起始位置
	// 直接分配一个S长度的栈，避免频繁的底层数组扩容
	stack := make([]byte, 0, len(S))
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			stack = append(stack, S[i])
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				// 当前原语的位置是S[start:i+1]
				result += S[start+1 : i]
				start = i + 1
			}
		}
	}

	return result
}

// 法二：不使用栈
// 只使用一个变量统计括号，当括号个数>=1，就添加到结果
// 空间O(1)，时间O(n)
func removeOuterParentheses2(S string) string {
	if S == "" {
		return S
	}
	result := make([]byte, 0, len(S))
	count := 0
	for i := 0; i < len(S); i++ {
		if S[i] == ')' { // 尝试消除最外层括号
			count--
		}
		if count >= 1 { // 最外层括号仍然在
			result = append(result, S[i])
		}
		if S[i] == '(' {
			count++
		}
	}
	return string(result)
}
