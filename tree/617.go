package main

import "github.com/davecgh/go-spew/spew"

// 617. 合并二叉树
// 给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。
// 你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，否则不为 NULL 的节点将直接作为新二叉树的节点。
// 注意: 合并必须从两个树的根节点开始。
// https://leetcode-cn.com/problems/merge-two-binary-trees/
func main() {
	spew.Dump(mergeTrees(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}, &TreeNode{
		Val: 4,
	}))
}
func mergeTrees(t1 *TreeNode, t2 *TreeNode) (t *TreeNode) {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t = &TreeNode{Val: t1.Val + t2.Val}
	t.Left = mergeTrees(t1.Left, t2.Left)
	t.Right = mergeTrees(t1.Right, t2.Right)
	return t
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
