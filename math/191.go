package main

import (
	"fmt"
)

// 191. 位1的个数
// 编写一个函数，输入是一个无符号整数，返回其二进制表达式中数字位数为 ‘1’ 的个数（也被称为汉明重量）。
// 提示：
//     请注意，在某些语言（如 Java）中，没有无符号整数类型。在这种情况下，输入和输出都将被指定为有符号整数类型，并且不应影响您的实现，因为无论整数是有符号的还是无符号的，其内部的二进制表示形式都是相同的。
//     在 Java 中，编译器使用二进制补码记法来表示有符号整数。因此，在上面的 示例 3 中，输入表示有符号整数 -3。
// https://leetcode-cn.com/problems/number-of-1-bits
func main() {
	fmt.Println(hammingWeight2(3)) // 2
}

// 法一
func hammingWeight(num uint32) (result int) {
	for num > 0 {
		if ((num - 1) ^ num) == 1 {
			result++
		}
		num >>= 1
	}
	return result
}

// 法二：另一种写法
// best
func hammingWeight2(num uint32) (result int) {
	for num != 0 {
		result++
		num &= (num - 1)
	}
	return result
}
