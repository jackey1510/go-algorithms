package main

import "fmt"

func main() {
	root := &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 6}}
	ans := kthSmallest(root, 3)

	fmt.Println(ans)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue[T any] []T

func (q *Queue[TreeNode]) poll() *TreeNode {
	if len(*q) == 0 {
		return nil
	}
	first := (*q)[0]
	*q = (*q)[1:]
	return &first
}

func (q *Queue[TreeNode]) push(node *TreeNode) {
	*q = append(*q, *node)
}

func kthSmallest(root *TreeNode, k int) int {
	result := -1

	dfs(root, &result, &k)
	return result
}

func dfs(node *TreeNode, result, k *int) {
	if node == nil {
		return
	}

	dfs(node.Left, result, k)
	*k--
	if *k == 0 {
		*result = node.Val
	}

	dfs(node.Right, result, k)

}
