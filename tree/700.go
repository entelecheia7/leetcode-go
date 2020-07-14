package main

import (
	"github.com/davecgh/go-spew/spew"
)

// 700. 二叉搜索树中的搜索
// 给定二叉搜索树（BST）的根节点和一个值。 你需要在BST中找到节点值等于给定值的节点。 返回以该节点为根的子树。 如果节点不存在，则返回 NULL。
// https://leetcode-cn.com/problems/search-in-a-binary-search-tree/
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

	spew.Dump(searchBST(tree, 5))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 法一：递归
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if val > root.Val {
		return searchBST(root.Right, val)
	}
	return searchBST(root.Left, val)
}

// 法二：循环
func searchBST2(root *TreeNode, val int) *TreeNode {
	for root != nil && root.Val != val {
		if val > root.Val {
			root = root.Right
		} else {
			root = root.Left
		}

	}
	return root
}
