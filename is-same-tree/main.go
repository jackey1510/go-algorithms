package main

func main() {

}

// TreeNode is
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return bfs(p, q)
}

func bfs(p *TreeNode, q *TreeNode) bool {

	if p == nil || q == nil {
		if p == q {
			return true
		}
		return false
	}

	if p.Val != q.Val {
		return false
	}
	return (bfs(p.Left, q.Left) && bfs(p.Right, q.Right))

}
