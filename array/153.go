package main

import (
	"fmt"
)

// 153. 寻找旋转排序数组中的最小值
// 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
// ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
// 请找出其中最小的元素。
// 你可以假设数组中不存在重复元素。
// https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/
func main() {
	fmt.Println(findMin2([]int{3, 4, 5, 1, 2}))
}

// 法一：二分查找
func findMin(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}
	left, right := 0, n-1
	min := nums[0]
	for left <= right {
		mid := left + ((right - left) >> 1)
		// 左侧有序，左侧的最小值已确定为nums[left]，试图向右侧寻找更小值
		if nums[left] <= nums[mid] {
			min = getMin(min, nums[left])
			left = mid + 1
		} else { // 右侧有序
			min = getMin(min, nums[mid])
			right = mid - 1
		}
	}
	return min
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 法二：二分查找的另一种写法
// 最小值永远在中点（考虑中点正好是发生了旋转的地方）以及发生了旋转的一侧
func findMin2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}
	left, right := 0, n-1
	for left < right {
		mid := left + ((right - left) >> 1)
		if nums[mid] > nums[right] { // 左侧有序，右侧无序
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}
