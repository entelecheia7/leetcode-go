package main

import "github.com/davecgh/go-spew/spew"

// 105. 从前序与中序遍历序列构造二叉树
// 根据一棵树的前序遍历与中序遍历构造二叉树。
// 注意:
// 你可以假设树中没有重复的元素。
// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
func main() {
	spew.Dump(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历为：中-左-右
// 中序遍历为：左-中-右
// 参考106.从中序与后序遍历序列构造二叉树 直接写出优化后的方案
func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(inorder)
	if n == 0 {
		return nil
	}
	inorderMap := make(map[int]int, n)
	for i := 0; i < n; i++ {
		inorderMap[inorder[i]] = i
	}
	rootIndex := 0
	return buildTreeHelper(preorder, &rootIndex, inorder, 0, n-1, inorderMap)
}
func buildTreeHelper(preorder []int, rootIndex *int, inorder []int, left, right int, inorderMap map[int]int) *TreeNode {
	if left > right {
		return nil
	}
	root := &TreeNode{
		Val: preorder[*rootIndex],
	}
	(*rootIndex)++
	root.Left = buildTreeHelper(preorder, rootIndex, inorder, left, inorderMap[root.Val]-1, inorderMap)
	root.Right = buildTreeHelper(preorder, rootIndex, inorder, inorderMap[root.Val]+1, right, inorderMap)
	return root
}
