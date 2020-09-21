package main

import "github.com/davecgh/go-spew/spew"

// 538. 把二叉搜索树转换为累加树
// 给定一个二叉搜索树（Binary Search Tree），把它转换成为累加树（Greater Tree)，使得每个节点的值是原来的节点值加上所有大于它的节点值之和。
// https://leetcode-cn.com/problems/convert-bst-to-greater-tree/
func main() {
	r := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	convertBST(r)
	spew.Dump(r)
}

// 一个节点应当加上它父节点和其父节点右子树的值
// 使用右-中-左的方式计算
func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	sum := 0
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root != nil {
			helper(root.Right)
			sum += root.Val
			root.Val = sum
			helper(root.Left)
		}
	}
	helper(root)
	return root
}

func helper(root *TreeNode) int {
	if root != nil {
		root.Val += helper(root.Right)
		helper(root.Left)
	}
	return 0
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
