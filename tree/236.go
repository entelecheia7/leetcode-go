package main

import "fmt"

// 236. 二叉树的最近公共祖先
// 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
// 说明:
//     所有节点的值都是唯一的。
//     p、q 为不同节点且均存在于给定的二叉树中。
// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
func main() {
	p := &TreeNode{Val: 5}
	q := &TreeNode{Val: 0}
	tree := &TreeNode{
		Val:  3,
		Left: p,
		Right: &TreeNode{
			Val:  1,
			Left: q,
		},
	}
	fmt.Println(lowestCommonAncestor(tree, p, q))
	fmt.Println(lowestCommonAncestor2(tree, p, q))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 法一：递归
// best
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 终结条件
	if root == nil {
		return nil
	} else if root == p || root == q {
		return root
	}
	// 处理当前层
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	// 如果left和right均不为空，说明p、q分别在root的两侧子树
	return root
}

// 法二：使用一个map保存节点的父节点
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q {
		return root
	}
	parent := make(map[int]*TreeNode)
	getParentNode(root, &parent)
	visited := make(map[int]bool)
	for p != nil {
		visited[p.Val] = true
		p = parent[p.Val]
	}
	for q != nil {
		if visited[q.Val] {
			return q
		}
		q = parent[q.Val]
	}
	return nil
}
func getParentNode(root *TreeNode, parent *map[int]*TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		(*parent)[root.Left.Val] = root
		getParentNode(root.Left, parent)
	}
	if root.Right != nil {
		(*parent)[root.Right.Val] = root
		getParentNode(root.Right, parent)
	}
}
