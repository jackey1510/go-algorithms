package main

import "fmt"

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}
	// root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	// root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	ans := diameterOfBinaryTree(root)
	fmt.Println(ans)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var diameter int

func diameterOfBinaryTree(root *TreeNode) int {
	diameter = 0
	return dfs(root)
}

func dfs(node *TreeNode) int {
	if node == nil {
		return 0
	}
	lengthLeft := dfs(node.Left)
	lengthRight := dfs(node.Right)
	diameter = max(diameter, lengthRight+lengthLeft+1)

	return max(lengthLeft, lengthRight) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
