package main

import (
	"fmt"
)

// 112. 路径总和
// 给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。
// 说明: 叶子节点是指没有子节点的节点。
// https://leetcode-cn.com/problems/path-sum/
func main() {
	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}
	fmt.Println(hasPathSum2(tree, 22))

	tree = &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
	}
	fmt.Println(hasPathSum2(tree, 1))

	fmt.Println(hasPathSum2(nil, 1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 法一：递归
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == sum
	}

	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}

// 法二：迭代，广度优先，使用一个队列
func hasPathSum2(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	queue := []*TreeNode{root}
	val := []int{root.Val}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		v := val[0]
		val = val[1:]
		if cur.Left == nil && cur.Right == nil {
			if v == sum {
				return true
			}
			continue
		}
		if cur.Left != nil {
			queue = append(queue, cur.Left)
			val = append(val, v+cur.Left.Val)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
			val = append(val, v+cur.Right.Val)
		}
	}

	return false
}
