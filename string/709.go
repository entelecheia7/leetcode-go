package main

import "fmt"

// 709. 转换成小写字母
// 实现函数 ToLowerCase()，该函数接收一个字符串参数 str，并将该字符串中的大写字母转换成小写字母，之后返回新的字符串。
// https://leetcode-cn.com/problems/to-lower-case/
func main() {
	fmt.Println(toLowerCase("He"))
}
func toLowerCase(str string) string {
	n := len(str)
	res := []byte(str)
	var diff byte = 'a' - 'A'
	for i := 0; i < n; i++ {
		if str[i] >= 'A' && str[i] <= 'Z' {
			res[i] = str[i] + diff
		}
	}
	return string(res)
}
