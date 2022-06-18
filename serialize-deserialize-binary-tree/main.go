package main

//https://leetcode.com/problems/serialize-and-deserialize-binary-tree/
import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	sec := Constructor()

	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}}, Right: &TreeNode{Val: 5}}}

	s := sec.serialize(root)
	fmt.Println(s)

	root2 := sec.deserialize(s)

	s2 := sec.serialize(root2)

	fmt.Println(s2)

	printTree(root2)
}

func printTree(root *TreeNode) {
	if root == nil {
		fmt.Println("null")
		return
	}
	fmt.Println(root.Val)
	printTree(root.Left)
	printTree(root.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	// LevelOrder [][]int
	serializeString strings.Builder
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	this.serializeString = strings.Builder{}
	traverse(root, this)

	return this.serializeString.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	split := strings.Split(data, ",")
	return deserializeHelper(&split)
}

func deserializeHelper(queue *[]string) *TreeNode {
	value := (*queue)[0]
	*queue = (*queue)[1:]
	if value == "null" {
		return nil
	}
	root := &TreeNode{Val: parseInt(value)}
	root.Left = deserializeHelper(queue)
	root.Right = deserializeHelper(queue)
	return root

}

func parseInt(s string) int {
	intVal, _ := strconv.ParseInt(s, 10, 32)
	return int(intVal)
}

func (this *Codec) appendNode(data string) {
	this.serializeString.WriteString(data)
	this.serializeString.WriteString(",")
}

func traverse(node *TreeNode, codec *Codec) {
	if node == nil {
		codec.appendNode("null")
		return
	}

	codec.appendNode(strconv.Itoa(node.Val))

	traverse(node.Left, codec)

	traverse(node.Right, codec)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
