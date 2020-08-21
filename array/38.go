package main

import (
	"fmt"
	"regexp"
)

// 38. 外观数列
// 给定一个正整数 n（1 ≤ n ≤ 30），输出外观数列的第 n 项。
// 注意：整数序列中的每一项将表示为一个字符串。
// 「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。前五项如下：
// 1.     1
// 2.     11
// 3.     21
// 4.     1211
// 5.     111221
// 第一项是数字 1
// 描述前一项，这个数是 1 即 “一个 1 ”，记作 11
// 描述前一项，这个数是 11 即 “两个 1 ” ，记作 21
// 描述前一项，这个数是 21 即 “一个 2 一个 1 ” ，记作 1211
// 描述前一项，这个数是 1211 即 “一个 1 一个 2 两个 1 ” ，记作 111221
// https://leetcode-cn.com/problems/count-and-say/
func main() {
	fmt.Println(countAndSay2(4)) // 1211
}

// 法一：暴力递推
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	prev := []byte{'1'}
	for i := 1; i < n; i++ {
		cur := []byte{}
		start := 0
		for start < len(prev) {
			tmp := start
			for start < len(prev)-1 && prev[start] == prev[start+1] {
				start++
			}
			cur = append(cur, byte('0'+(start-tmp+1)))
			cur = append(cur, prev[start])
			start++
		}
		prev = cur

	}
	return string(prev)
}

// 法二：使用正则表达式
//  (\d)\1* 可以匹配连续相同数字，通过将n-1 的结果进行匹配替换获得n
// 对于Go语言不太方便，因为Go不支持反向引用
