package main

import (
	"fmt"
	"strconv"
)

// 412. Fizz Buzz
// 写一个程序，输出从 1 到 n 数字的字符串表示。
// 1. 如果 n 是3的倍数，输出“Fizz”；
// 2. 如果 n 是5的倍数，输出“Buzz”；
// 3.如果 n 同时是3和5的倍数，输出 “FizzBuzz”。
// https://leetcode-cn.com/problems/fizz-buzz
func main() {
	fmt.Println(fizzBuzz2(15))
}

// 法一：暴力法
func fizzBuzz(n int) []string {
	result := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			result = append(result, "FizzBuzz")
		} else if i%5 == 0 {
			result = append(result, "Buzz")
		} else if i%3 == 0 {
			result = append(result, "Fizz")
		} else {
			// result = append(result, fmt.Sprintf("%d", i))
			// strconv的空间复杂度略低
			result = append(result, strconv.Itoa(i))
		}
	}

	return result
}

// 法二：暴力的另一种写法，避免了除法操作
func fizzBuzz2(n int) []string {
	result := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		result = append(result, strconv.Itoa(i))
	}
	for i := 2; i < n; i += 3 {
		result[i] = "Fizz"
	}
	for i := 4; i < n; i += 5 {
		result[i] = "Buzz"
	}
	for i := 14; i < n; i += 15 {
		result[i] = "FizzBuzz"
	}

	return result
}

// 法三：在写法上对法一进行优化，但效率有所降低
// 如果是一个map遍历有序的语言，可以把这种映射关系放到map，以实现更广泛的映射关系
func fizzBuzz3(n int) []string {
	result := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		var char string
		if i%3 == 0 {
			char += "Fizz"
		}
		if i%5 == 0 {
			char += "Buzz"
		}
		if len(char) == 0 {
			char = fmt.Sprintf("%d", i)
		}
		result = append(result, char)
	}

	return result
}
