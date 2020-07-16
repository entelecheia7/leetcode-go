package main

import (
	"fmt"
)

// 96. 不同的二叉搜索树
// 给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？
// https://leetcode-cn.com/problems/unique-binary-search-trees/
func main() {
	fmt.Println(numTrees(3))  // 5
	fmt.Println(numTrees2(3)) // 5
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 类似题目注意是否包含空树
// 法一：动态规划
// n个整数组成的BST个数=以i为根节点的BST个数相加（1<=i<=n）
// 以i为根节点的BST，左子树是i-1个节点，右子树是n-i个节点
// 则根为i的BST集合是左子树集合和右子树集合的笛卡尔积
// O(n^2)
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ { // 表示拥有n个节点的BST数目统计
		for j := 1; j <= i; j++ { // j表示根节点i的可选范围，统计以i为根节点的可能性之和
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

// 法二：卡特兰通项公式
// 根据法一，举例如下：
// Taking 1~n as root respectively:
// 1 as root: # of trees = F(0) * F(n-1)  // F(0) == 1
// 2 as root: # of trees = F(1) * F(n-2)
// 3 as root: # of trees = F(2) * F(n-3)
// ...
// n-1 as root: # of trees = F(n-2) * F(1)
// n as root:   # of trees = F(n-1) * F(0)
// so，F(n) = F(0) * F(n-1) + F(1) * F(n-2) + F(2) * F(n-3) + ... + F(n-2) * F(1) + F(n-1) * F(0)
// 符合卡特兰公式：
// 令h(0)=1,h(1)=1。
// 卡塔兰数的递推式：h(n) = h(0)*h(n-1) + h(1)*h(n-2) + ... + h(n-1)h(0)（n>=2）
// h(n) = h(n-1)*(4*n-2) / (n+1)
func numTrees2(n int) int {
	hi := 1 // n=1
	for i := 2; i <= n; i++ {
		hi = hi * (4*i - 2) / (i + 1)
	}
	return hi
}
