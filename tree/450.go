package main

import (
	"github.com/davecgh/go-spew/spew"
)

// 450. 删除二叉搜索树中的节点
// 给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。
// 一般来说，删除节点可分为两个步骤：
//     首先找到需要删除的节点；
//     如果找到了，删除它。
// 说明： 要求算法时间复杂度为 O(h)，h 为树的高度。
// https://leetcode-cn.com/problems/delete-node-in-a-bst/
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

	spew.Dump(deleteNode(tree, 2))
	spew.Dump(deleteNode(&TreeNode{Val: 0}, 0))
	tree = &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 4,
		},
	}
	spew.Dump(deleteNode(tree, 3))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 法一：递归
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else { // root 是待删除节点
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			// root 拥有两个子节点
			root.Val := getReplaceNodeFromLeft(root).Val
			root.Left = deleteNode(root.Left, root.Val)
		}
	}
	return root
}

// 替代节点可以是左子树的最右节点或者右子树的最左节点
// 这里使用的是左子树的最右节点
// 返回替代节点
func getReplaceNodeFromLeft(root *TreeNode) (node *TreeNode) {
	node = root.Left
	for node.Right != nil {
		node = node.Right
	}
	return node
}
// 法二：删除节点部分的另一种写法
func deleteNode2(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else { // root 是待删除节点
		if root.Left == nil && root.Right == nil {
			root = nil
		} else if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			// root 拥有两个子节点
			replaceNode := getReplaceNodeFromLeft(root)
			replaceNode.Right = root.Right
			return root.Left
		}
	}
	return root
}