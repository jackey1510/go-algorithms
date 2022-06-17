package main

func main() {
	isBalanced(nil)
}

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//https://leetcode.com/problems/balanced-binary-tree/
func isBalanced(root *TreeNode) bool {
	return bfs(root, 0) != -1
}

func bfs(node *TreeNode, depth int) int {
	if node == nil {
		return depth
	}

	l := bfs(node.Left, depth+1)
	r := bfs(node.Right, depth+1)

	if l == -1 || r == -1 || abs(l-r) > 1 {
		return -1
	}

	return max(l, r)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}
