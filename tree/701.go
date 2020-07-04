package main

import (
	"github.com/davecgh/go-spew/spew"
)

// 701. 二叉搜索树中的插入操作
// 给定二叉搜索树（BST）的根节点和要插入树中的值，将值插入二叉搜索树。 返回插入后二叉搜索树的根节点。 保证原始二叉搜索树中不存在新值。
// 注意，可能存在多种有效的插入方式，只要树在插入后仍保持为二叉搜索树即可。 你可以返回任意有效的结果。
// https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/
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

	spew.Dump(insertIntoBST(tree, 5))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉搜索树的插入时间复杂度是O(logn)

// 法一：递归
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val > root.Val {
		root.Right = insertIntoBST(root.Right, val)
	} else {
		root.Left = insertIntoBST(root.Left, val)
	}

	return root
}

// 法二：循环
func insertIntoBST2(root *TreeNode, val int) *TreeNode {
	newNode := &TreeNode{Val: val}
	if root == nil {
		return newNode
	}
	p := root
	for p != nil {
		if val > p.Val {
			if p.Right == nil {
				p.Right = newNode
				break
			} else {
				p = p.Right
			}
		} else {
			if p.Left == nil {
				p.Left = newNode
				break
			} else {
				p = p.Left
			}
		}
	}

	return root
}
