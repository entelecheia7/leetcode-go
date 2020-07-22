package main

import (
	"github.com/davecgh/go-spew/spew"
)

// 95. 不同的二叉搜索树 II
// 给定一个整数 n，生成所有由 1 ... n 为节点所组成的 二叉搜索树 。
// 提示：
// 0 <= n <= 8
// https://leetcode-cn.com/problems/unique-binary-search-trees-ii/
func main() {
	// spew.Dump(generateTrees(3))
	spew.Dump(generateTrees2(3))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 法一：递归
// 从1-n每个数字都可以做根节点
// 假设以i为根节点（1<=i<=n），[1,i)为左子树，[i+1,n]为右子树
func generateTrees(n int) []*TreeNode {
	//特殊情况
	if n == 0 {
		return nil
	} else if n == 1 {
		return []*TreeNode{
			{Val: 1},
		}
	}

	return generateTreesHelper(1, n)
}

func generateTreesHelper(left, right int) (result []*TreeNode) {
	if left > right {
		return nil
	} else if left == right {
		return []*TreeNode{
			{Val: left},
		}
	}
	for i := left; i <= right; i++ {
		left := generateTreesHelper(left, i-1)
		right := generateTreesHelper(i+1, right)
		if len(left) == 0 {
			for _, r := range right {
				result = append(result, &TreeNode{
					Val:   i,
					Right: r,
				})
			}
			continue
		}
		if len(right) == 0 {
			for _, l := range left {
				result = append(result, &TreeNode{
					Val:  i,
					Left: l,
				})
			}
			continue
		}
		for _, l := range left {
			for _, r := range right {
				result = append(result, &TreeNode{
					Val:   i,
					Left:  l,
					Right: r,
				})
			}
		}
	}
	return result
}

// 法二：动态规划
// 求解f(n)相当于在 f(n-1)里添加节点n，n可以是根节点或f(n-1)右子树中的节点
// 但是将节点n插入f(n-1)的右子树需要考虑在每一层插入的情况
