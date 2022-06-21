package main

//https://leetcode.com/problems/binary-tree-right-side-view/submissions/
func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	res := []int{}
	queue := [][]*TreeNode{{root}}
	levels := [][]int{}

	level := 0

	for len(queue) > level {
		if len(queue[level]) == 0 {
			level++
			continue
		}
		node := queue[level][0]
		queue[level] = queue[level][1:]
		if node == nil {
			continue
		}
		if len(levels) == level {
			levels = append(levels, []int{})
		}
		levels[level] = append(levels[level], node.Val)
		if len(queue) == level {
			queue = append(queue, []*TreeNode{})
		}
		queue[level+1] = append(queue[level+1], node.Right, node.Left)
	}
	for _, lev := range levels {
		res = append(res, lev[0])
	}
	return res
}
