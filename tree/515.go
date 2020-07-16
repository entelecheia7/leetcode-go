package main

import (
	"fmt"
	"math"
)

// 515. 在每个树行中找最大值
// 您需要在二叉树的每一行中找到最大的值。
// https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row/#/description
func main() {
	tree := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 7,
		},
	}
	fmt.Println(largestValues(tree))
	fmt.Println(largestValues2(tree)) //best
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 法一：BFS，相当于层序遍历
func largestValues(root *TreeNode) (result []int) {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		max := math.MinInt64
		for i := 0; i < size; i++ {
			max = getMax(max, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[size:]
		result = append(result, max)
	}
	return result
}
func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 法二：DFS
// 不需要额外的空间，空间复杂度优于法一
func largestValues2(root *TreeNode) (result []int) {
	if root == nil {
		return nil
	}
	largestValuesDFSHelper(root, 0, &result)
	return result
}
func largestValuesDFSHelper(root *TreeNode, level int, result *[]int) {
	if len(*result) == level {
		*result = append(*result, math.MinInt64)
	}
	(*result)[level] = getMax((*result)[level], root.Val)
	if root.Left != nil {
		largestValuesDFSHelper(root.Left, level+1, result)
	}
	if root.Right != nil {
		largestValuesDFSHelper(root.Right, level+1, result)
	}
}

// 法三：也可以使用分治，略
