package main

func main() {

}

// https://leetcode.com/problems/binary-tree-level-order-traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	bfs(root, &res, 0)
	return res
}

func bfs(root *TreeNode, res *[][]int, depth int) {
	if root == nil {
		return
	}
	l := len(*res)
	if l == depth {
		*res = append(*res, []int{})
	}

	(*res)[depth] = append((*res)[depth], root.Val)

	bfs(root.Left, res, depth+1)

	bfs(root.Right, res, depth+1)
}
