package main

import (
	// "fmt"
	"github.com/davecgh/go-spew/spew"
)

// 43. 字符串相乘
// 给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
// 说明：
//     num1 和 num2 的长度小于110。
//     num1 和 num2 只包含数字 0-9。
//     num1 和 num2 均不以零开头，除非是数字 0 本身。
//     不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
// https://leetcode-cn.com/problems/multiply-strings
func main() {
	spew.Dump(multiply2("98", "9")) // 882
	spew.Dump(multiply2("6", "501"))
	spew.Dump(multiply2("2", "3"))
	spew.Dump(multiply2("123", "456")) // 56088
}

// 法一：暴力，按照竖式乘法计算，时间复杂度高
func multiply(num1 string, num2 string) (ans string) {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	n1, n2 := len(num1), len(num2)
	if n1 > n2 {
		n1, n2 = n2, n1
		num1, num2 = num2, num1
	}
	ans = "0"
	for i := n1 - 1; i >= 0; i-- {
		// 使用 num1[i] * num2的各个位
		if num1[i] == '0' {
			continue
		}
		var tmp []byte
		for k := 0; k < n1-1-i; k++ {
			tmp = append(tmp, '0')
		}
		// 右侧补0
		cur := int(num1[i] - '0')
		carry := 0
		for j := n2 - 1; j >= 0; j-- {
			mul := cur*int(num2[j]-'0') + carry
			carry = mul / 10
			tmp = append([]byte{byte(mul%10 + '0')}, tmp...)
		}
		if carry > 0 {
			ans = add(ans, string([]byte{byte(carry + '0')})+string(tmp))
		} else {
			ans = add(ans, string(tmp))
		}
	}
	return string(ans)
}

// n2>=0，n1>=0
func add(n1 string, n2 string) string {
	if n1 == "0" {
		return n2
	}
	if n2 == "0" {
		return n1
	}
	i, j := len(n1)-1, len(n2)-1
	carry := 0
	k := getMax(len(n1), len(n2)) - 1
	result := make([]byte, k+1)
	for i >= 0 || j >= 0 {
		sum := carry
		if i >= 0 {
			sum += int(n1[i] - '0')
		}
		if j >= 0 {
			sum += int(n2[j] - '0')
		}
		carry = sum / 10
		result[k] = byte(sum%10 + '0')
		i--
		j--
		k--
	}
	if carry == 0 {
		return string(result)
	}
	return string([]byte{byte(carry + '0')}) + string(result)
}

func getMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// 法二：竖式乘法优化
// 设 num1 长度为 n1，num2 长度为 n2。最多需要 n1+n2 位
// 先将个位乘法计算出来，个位保存在 i+j+1 位，十位保存在 i+j 位置func multiply2(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	n1, n2 := len(num1), len(num2)
	res := make([]int, n1+n2)
	for i := n1 - 1; i >= 0; i-- {
		// 使用 num1[i] * num2的各个位
		cur := int(num1[i] - '0')
		for j := n2 - 1; j >= 0; j-- {
			mul := cur*int(num2[j]-'0') + res[i+j+1]
			res[i+j+1] = mul % 10
			res[i+j] += mul / 10
		}
	}
	result := make([]byte, n1+n2)
	for k, num := range res {
		result[k] = byte(num + '0')
	}
	if result[0] == '0' {
		return string(result[1:])
	}
	return string(result)
}

