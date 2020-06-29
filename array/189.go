package main

import (
	"fmt"
)

// 189. 旋转数组
// 给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
// 说明:
// 尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
// 要求使用空间复杂度为 O(1) 的 原地 算法。
// https://leetcode-cn.com/problems/rotate-array/
func main() {
	var nums []int
	nums = []int{1, 2, 3, 4, 5, 6, 7}
	rotate2(nums, 3)
	fmt.Println(nums)
	nums = []int{1, 2, 3, 4, 5, 6, 7}
	rotate3(nums, 3)
	fmt.Println(nums)
}

// 法一：使用一个额外的数组，空间O(n)，时间O(n)，非原地算法，略

// 法二：每次将元素向右旋转一位，旋转k次，时间O(k*n)，空间O(1)，原地算法
func rotate2(nums []int, k int) {
	n := len(nums)
	if n < 2 {
		return
	}
	tmp := nums[0]
	for k > 0 {
		for i := 0; i < n; i++ {
			if i == n-1 {
				nums[0] = tmp
			} else {
				tmp, nums[i+1] = nums[i+1], tmp
			}
		}
		k--
	}
}

// 法三：三次数组反转。
// 旋转k次，数组末尾的 k%n 个元素会移动到数组头部，其余元素右移
// 首先将整个数组反转，然后反转前k个元素，再反转剩下的元素
// 空间O(1)，时间O(n)，原地算法
func rotate3(nums []int, k int) {
	n := len(nums)
	k %= n
	if n < 2 || k == 0 {
		return
	}
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)

}
func reverse(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

// 法四：环状替代
// 空间O(1)，时间O(n)，原地算法
func rotate4(nums []int, k int) {
	n := len(nums)
	k %= n
	if n < 2 || k == 0 {
		return
	}
	count := 0
	for i := 0; count < n; i++ {
		cur, val := i, nums[i]
	C:
		{
			next := (cur + k) % n
			tmp := nums[next]
			nums[next] = val
			cur, val = next, tmp
			count++
			if cur != i {
				goto C
			}
		}
	}
}
