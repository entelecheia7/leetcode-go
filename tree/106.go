package main

import "github.com/davecgh/go-spew/spew"

// 106. 从中序与后序遍历序列构造二叉树
// 根据一棵树的中序遍历与后序遍历构造二叉树。
// 注意:
// 你可以假设树中没有重复的元素。
// 例如，给出
// 中序遍历 inorder = [9,3,15,20,7]
// 后序遍历 postorder = [9,15,7,20,3]
// 返回如下的二叉树：
//     3
//    / \
//   9  20
//     /  \
//    15   7
// https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
func main() {
	spew.Dump(buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}))
	spew.Dump(buildTree2([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中序遍历的顺序是：左-中-右
// 后序遍历的顺序是：左-右-中
// 法一：递归
// 由后序遍历可得根节点，在中序遍历的结果中找到根节点下标 i
// 则 inorder[0:i] 为左子树，inorder[i+1:] 是右子树
// 递归处理
func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(postorder)
	if n == 0 {
		return nil
	}
	root := &TreeNode{
		Val: postorder[n-1],
	}
	if n == 1 {
		return root
	}

	i := 0
	for ; i < n; i++ {
		if inorder[i] == root.Val {
			break
		}
	}
	// inorder[0:i]为左子树，inorder[i+1:]是右子树
	root.Left = buildTree(inorder[:i], postorder[:i])
	root.Right = buildTree(inorder[i+1:], postorder[i:n-1])

	return root
}

// 法二：对法一进行优化
// 使用一个map去保存中序数组的下标和值的映射关系，省去遍历查找
// 使用数组下标来保存本次函数处理的数值范围，省去切片的创建
func buildTree2(inorder []int, postorder []int) *TreeNode {
	n := len(postorder)
	if n == 0 {
		return nil
	}
	root := &TreeNode{
		Val: postorder[n-1],
	}
	if n == 1 {
		return root
	}
	m := make(map[int]int, n) // val=>index
	for i := 0; i < n; i++ {
		m[inorder[i]] = i
	}
	return buildTreeHelper(inorder, 0, n-1, postorder, n-1, m)
}

// left 和 right 是中序数组本次函数需要处理的下标范围
func buildTreeHelper(inorder []int, left, right int, postorder []int, postIndex int, m map[int]int) *TreeNode {
	if left > right {
		return nil
	}
	root := &TreeNode{
		Val: postorder[postIndex],
	}
	inIndex := m[root.Val]
	root.Right = buildTreeHelper(inorder, inIndex+1, right, postorder, postIndex-1, m)
	// 后序遍历的顺序是左-右-中，因此左子树的根节点下标是：父节点下标-1-右子树长度
	root.Left = buildTreeHelper(inorder, left, inIndex-1, postorder, postIndex-1-right+inIndex, m)
	return root
}
