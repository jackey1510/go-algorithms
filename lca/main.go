package main

import (
	"fmt"
)

//https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
func main() {
	left, right := &TreeNode{Val: 2}, &TreeNode{Val: 8}
	root := &TreeNode{Val: 6, Left: left, Right: right}
	res := lowestCommonAncestor(root, left, right)

	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var min *TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if min == nil {
		min = root
	}

	min = minAncestor(min, root)

	if q.Val < p.Val {
		p, q = q, p
	}

	if p.Val < min.Val && q.Val > min.Val {
		return min
	}

	if q.Val < min.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	return lowestCommonAncestor(root.Right, p, q)
}

func minAncestor(a, b *TreeNode) *TreeNode {
	if b == nil {
		return a
	}
	if a.Val < b.Val {
		return a
	}
	return b
}
