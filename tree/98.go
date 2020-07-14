package main

import (
	"fmt"
	"math"
)

// 98. 验证二叉搜索树
// 给定一个二叉树，判断其是否是一个有效的二叉搜索树。
// 假设一个二叉搜索树具有如下特征：
//     节点的左子树只包含小于当前节点的数。
//     节点的右子树只包含大于当前节点的数。
//     所有左子树和右子树自身必须也是二叉搜索树。
// https://leetcode-cn.com/problems/validate-binary-search-tree/
func main() {
	tree1 := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	}
	tree2 := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}},
	}
	fmt.Println(isValidBST(tree1))
	fmt.Println(isValidBST(tree2))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 法一：递归
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return validateBSTHelper(root.Left, root.Val, math.MinInt64) && validateBSTHelper(root.Right, math.MaxInt64, root.Val)
}
func validateBSTHelper(root *TreeNode, maxVal, minVal int) bool {
	if root == nil {
		return true
	}
	if root.Val >= maxVal || root.Val <= minVal {
		return false
	}
	return validateBSTHelper(root.Left, root.Val, minVal) && validateBSTHelper(root.Right, maxVal, root.Val)
}

// 法二：中序遍历，是否升序
func isValidBST2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	stack := []*TreeNode{}
	cur := root
	lastVal := math.MinInt64
	for len(stack) > 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur.Val <= lastVal {
			return false
		}
		lastVal = cur.Val
		cur = cur.Right
	}
	return true
}
