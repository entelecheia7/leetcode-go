package main

import (
	"fmt"
)

// 1122. 数组的相对排序
// 给你两个数组，arr1 和 arr2，
//     arr2 中的元素各不相同
//     arr2 中的每个元素都出现在 arr1 中
// 对 arr1 中的元素进行排序，使 arr1 中项的相对顺序和 arr2 中的相对顺序相同。未在 arr2 中出现过的元素需要按照升序放在 arr1 的末尾。
// 提示：
//     arr1.length, arr2.length <= 1000
//     0 <= arr1[i], arr2[i] <= 1000
//     arr2 中的元素 arr2[i] 各不相同
//     arr2 中的每个元素 arr2[i] 都出现在 arr1 中
// https://leetcode-cn.com/problems/relative-sort-array/
func main() {
	fmt.Println(relativeSortArray([]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6})) // [2,2,2,1,4,3,3,9,6,7,19]
}

// 计数排序
// 由于数据范围不大，可以采用计数排序
func relativeSortArray(arr1 []int, arr2 []int) (result []int) {
	a1 := make([]int, 1001)
	for _, v := range arr1 {
		a1[v]++
	}
	for _, v := range arr2 {
		for a1[v] > 0 {
			result = append(result, v)
			a1[v]--
		}
	}
	for num, count := range a1 {
		for count > 0 {
			result = append(result, num)
			count--
		}
	}
	return result
}
