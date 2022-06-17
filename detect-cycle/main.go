package main

func main() {

}

// https://leetcode.com/problems/linked-list-cycle/
// ListNode is
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	return cycle(head, head.Next)
}

func cycle(node1 *ListNode, node2 *ListNode) bool {
	if node1 == nil || node2 == nil || node2.Next == nil {
		return false
	}
	if node1 == node2 {
		return true
	}

	return cycle(node1.Next, node2.Next.Next)

}
