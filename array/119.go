package main

import "fmt"

// 119. 杨辉三角 II
// 给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。
// 在杨辉三角中，每个数是它左上方和右上方的数的和。
// https://leetcode-cn.com/problems/pascals-triangle-ii/
func main() {
	fmt.Println(getRow(3)) // [1 3 3 1]
}

// 递推
func getRow(rowIndex int) []int {
	cur := []int{1}
	if rowIndex == 0 {
		return cur
	}
	for i := 1; i <= rowIndex; i++ {
		cur = append(cur, 0)
		for j := i; j > 0; j-- {
			cur[j] += cur[j-1]
		}
	}
	return cur
}
