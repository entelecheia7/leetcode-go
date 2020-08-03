package main

import (
	"fmt"
)

// 415. 字符串相加
// 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。
// 注意：
//     num1 和num2 的长度都小于 5100.
//     num1 和num2 都只包含数字 0-9.
//     num1 和num2 都不包含任何前导零。
//     你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。
// https://leetcode-cn.com/problems/add-strings/
func main() {
	fmt.Println(addStrings("123", "59"))
	fmt.Println(addStrings("1", "9"))
}

func addStrings(num1 string, num2 string) string {
	var result []byte
	i, j, k := len(num1)-1, len(num2)-1, 0
	if i > j {
		result = make([]byte, len(num1))
		k = i
	} else {
		result = make([]byte, len(num2))
		k = j
	}
	var tmp, carry byte
	for i >= 0 || j >= 0 {
		tmp = carry
		if i >= 0 {
			tmp += num1[i] - '0'
			i--
		}
		if j >= 0 {
			tmp += num2[j] - '0'
			j--
		}
		result[k] = tmp%10 + '0'
		k--
		carry = tmp / 10
	}
	if carry == 1 {
		result = append([]byte{'1'}, result...)
	}

	return string(result)
}
