package main

import "fmt"

// 59. 螺旋矩阵 II
// 给定一个正整数 n，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。
// 示例:
// 输入: 3
// 输出:
// [
//  [ 1, 2, 3 ],
//  [ 8, 9, 4 ],
//  [ 7, 6, 5 ]
// ]
// https://leetcode-cn.com/problems/spiral-matrix-ii/
func main() {
	fmt.Println(generateMatrix(3))
	fmt.Println(generateMatrix2(3))
}

// 法一：暴力法
func generateMatrix(n int) (result [][]int) {
	result = make([][]int, n)
	for k := range result {
		result[k] = make([]int, n)
	}
	max := n * n
	left, right, top, bottom := 0, n-1, 0, n-1
	num := 1
	for num <= max {
		// left=>right
		for i := left; i <= right; i, num = i+1, num+1 {
			result[top][i] = num
		}
		top++
		// top => bottom
		for i := top; i <= bottom; i, num = i+1, num+1 {
			result[i][right] = num
		}
		right--
		// right=>left
		for i := right; i >= left; i, num = i-1, num+1 {
			result[bottom][i] = num
		}
		bottom--
		// bottom=>top
		for i := bottom; i >= top; i, num = i-1, num+1 {
			result[i][left] = num
		}
		left++
	}

	return result
}

// 法二：归纳法
func generateMatrix2(n int) (result [][]int) {
	result = make([][]int, n)
	for k := range result {
		result[k] = make([]int, n)
	}
	max := n * n
	num := 1
	i, j := 0, 0
	di, dj := 0, 1 // result[i][j]距下一个点的距离
	for num <= max {
		result[i][j] = num
		nexti, nextj := (i+di)%n, (j+dj)%n
		if nextj < 0 { // handle bottom=>top index out of range
			nextj += n
		}
		// left=>right di:0,dj:1
		// top => bottom di:1,dj:0
		// right=>left di:0,dj:-1
		// bottom=>top di:-1,dj:0
		if result[nexti][nextj] > 0 { // 说明本行或本列已经处理完毕
			di, dj = dj, -di
		}
		i += di
		j += dj
		num++
	}
	return result
}
