package main

import "fmt"

// 118. 杨辉三角
// 给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
// 在杨辉三角中，每个数是它左上方和右上方的数的和。
// https://leetcode-cn.com/problems/pascals-triangle/
func main() {
	fmt.Println(generate(5))
}

func generate(numRows int) (result [][]int) {
	if numRows == 0 {
		return nil
	}
	result = make([][]int, numRows)
	result[0] = []int{1}
	for i := 1; i < numRows; i++ {
		result[i] = make([]int, i+1)
		result[i][0], result[i][i] = 1, 1
		for j := 1; j < i; j++ {
			result[i][j] = result[i-1][j] + result[i-1][j-1]
		}
	}

	return result
}
