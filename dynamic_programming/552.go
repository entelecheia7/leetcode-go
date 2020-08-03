package main

import (
	"fmt"
)

// 552. 学生出勤记录 II
// 给定一个正整数 n，返回长度为 n 的所有可被视为可奖励的出勤记录的数量。 答案可能非常大，你只需返回结果mod 109 + 7的值。
// 学生出勤记录是只包含以下三个字符的字符串：
//     'A' : Absent，缺勤
//     'L' : Late，迟到
//     'P' : Present，到场
// 如果记录不包含多于一个'A'（缺勤）或超过两个连续的'L'（迟到），则该记录被视为可奖励的。
// 注意：n 的值不会超过100000。
// https://leetcode-cn.com/problems/student-attendance-record-ii/
func main() {
	fmt.Println(checkRecord(1)) // 3
	fmt.Println(checkRecord(2)) // 8

	// fmt.Println(checkRecord(100))
}

// 法一：回溯。时间复杂度过高，无法AC，略

// 法二：动态规划
// 将可能性分为6种：
// a 不含A的LL结尾，可添加A P
// b 不含A的L结尾，倒数第二位不是L，可添加A L P
// c 不含A的非L结尾，可添加A L P
// d 含A的LL结尾，可添加P
// e 含A的L结尾，倒数第二位不为L，可添加 L P
// f 含A的非L结尾，可添加 L P
func checkRecord(n int) int {
	// 初始化n=1
	a, b, c, d, e, f := 0, 1, 1, 0, 0, 1
	// i 本该从 2 开始
	// 但由于最后的结果是(a + b + c + d + e + f) % 1000000007
	// f恰好是a-f的和，所以直接多循环一遍
	for i := 1; i <= n; i++ {
		a, b, c = b, c, (a+b+c)%1000000007
		d, e, f = e, f, (d+e+f+c)%1000000007
	}

	return f
}
