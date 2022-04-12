package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	ans := math.MinInt
	bfs(root, &ans)
	return ans
}

func bfs(node *TreeNode, ans *int) int {
	if node == nil {
		return 0
	}
	leftSum := maxNum(bfs(node.Left, ans), 0)
	rightSum := maxNum(bfs(node.Right, ans), 0)

	sum := node.Val + leftSum + rightSum

	*ans = maxNum(sum, *ans)

	return maxNum(leftSum, rightSum) + node.Val

}

func maxNum(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
