package main

func main() {

}

// ListNode is
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {

	node := head
	node2 := node.Next

	for node2 != nil && node2.Next != nil {
		if node.Val == node2.Val {
			return true
		}
		node = node.Next
		node2 = node2.Next
		if node2.Next == nil {
			return true
		}
		node2 = node2.Next
	}

	return true
}
