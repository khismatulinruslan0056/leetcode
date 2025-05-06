package main

/*
Given the root of a binary tree, return all root-to-leaf paths in any order.

A leaf is a node with no children.



Example 1:


Input: root = [1,2,3,null,5]
Output: ["1->2->5","1->3"]
Example 2:

Input: root = [1]
Output: ["1"]


Constraints:

The number of nodes in the tree is in the range [1, 100].
-100 <= Node.val <= 100 */

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		strVal := strconv.Itoa(root.Val)
		return []string{strVal}
	}
	left := binaryTreePaths(root.Left)
	right := binaryTreePaths(root.Right)
	for i := 0; i < len(left); i++ {
		left[i] = strconv.Itoa(root.Val) + "->" + left[i]
	}
	for i := 0; i < len(right); i++ {
		right[i] = strconv.Itoa(root.Val) + "->" + right[i]
	}
	return append(left, right...)
}

func main() {
	p := []interface{}{1, 2, 3, nil, 5}
	pTree := NewTree(p)
	fmt.Println(binaryTreePaths(pTree))
	p = []interface{}{1}
	pTree = NewTree(p)
	fmt.Println(binaryTreePaths(pTree))

}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val, Left: nil, Right: nil}
}

func NewTree(data []interface{}) *TreeNode {
	if len(data) == 0 || data[0] == nil {
		return nil
	}
	rootTreeNode := &TreeNode{Val: data[0].(int)}
	queue := []*TreeNode{rootTreeNode}
	i := 1
	for len(queue) > 0 && i < len(data) {
		current := queue[0]
		queue = queue[1:]
		if i < len(data) && data[i] != nil {
			current.Left = &TreeNode{Val: data[i].(int)}
			queue = append(queue, current.Left)
		}
		i++
		if i < len(data) && data[i] != nil {
			current.Right = &TreeNode{Val: data[i].(int)}
			queue = append(queue, current.Right)
		}
		i++

	}
	return rootTreeNode
}

func (t *TreeNode) InOrder2() {
	current := t
	stack := make([]*TreeNode, 0)
	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Print(current.Val, " ")
		current = current.Right
	}

}

func InOrder(t *TreeNode) {
	if t != nil {
		InOrder(t.Left)
		fmt.Print(t.Val, " ")
		InOrder(t.Right)
	}
}
