package main

import "fmt"

// 74. 搜索二维矩阵
// 编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
// 每行中的整数从左到右按升序排列。
// 每行的第一个整数大于前一行的最后一个整数。
// https://leetcode-cn.com/problems/search-a-2d-matrix/
func main() {
	fmt.Println(searchMatrix([][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}, 3))
}

// 二分查找
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 || matrix[0][0] > target {
		return false
	}
	m := len(matrix)
	n := len(matrix[0])
	if matrix[m-1][n-1] < target {
		return false
	}
	left, right := 0, m*n-1
	for left <= right {
		mid := left + ((right - left) >> 1)
		midVal := matrix[mid/n][mid%n]
		if midVal == target {
			return true
		} else if midVal < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
