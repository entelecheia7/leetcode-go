package main

import (
	"fmt"
)

// 104. 二叉树的最大深度
// 给定一个二叉树，找出其最大深度。
// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
// 说明: 叶子节点是指没有子节点的节点。
// https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
func main() {
	tree := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}

	fmt.Println(maxDepth(tree))
	fmt.Println(maxDepth2(tree))
}

// 法一：递归
// 时间O(n),空间O(logn)
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	return getMax(leftDepth, rightDepth) + 1
}
func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 法二：迭代
type MyNode struct {
	*TreeNode
	level int
}

func maxDepth2(root *TreeNode) (maxDepth int) {
	if root == nil {
		return 0
	}
	stack := []MyNode{{root, 1}}
	for len(stack) > 0 {
		last := stack[len(stack)-1]
		maxDepth = getMax(maxDepth, last.level)
		stack = stack[:len(stack)-1]
		if last.Left != nil {
			stack = append(stack, MyNode{last.Left, last.level + 1})
		}
		if last.Right != nil {
			stack = append(stack, MyNode{last.Right, last.level + 1})
		}
	}
	return
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
