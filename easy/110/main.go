package main

/*Given a binary tree, determine if it is height-balanced.



Example 1:


Input: root = [3,9,20,null,null,15,7]
Output: true
Example 2:


Input: root = [1,2,2,3,3,null,null,4,4]
Output: false
Example 3:

Input: root = []
Output: true*/

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftDepth := Depth(root.Left)
	rightDepth := Depth(root.Right)
	delta := leftDepth - rightDepth
	if rightDepth > leftDepth {
		delta = rightDepth - leftDepth
	}
	if delta > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func Depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := Depth(root.Left)
	rightDepth := Depth(root.Right)
	return 1 + max(leftDepth, rightDepth)
}

func main() {
	p := []interface{}{3, 9, 20, nil, nil, 15, 7}
	pTree := NewTree(p)
	fmt.Println(isBalanced(pTree))

	p = []interface{}{1, 2, 2, 3, 3, nil, nil, 4, 4}
	pTree = NewTree(p)
	fmt.Println(isBalanced(pTree))

	p = []interface{}{}
	pTree = NewTree(p)
	fmt.Println(isBalanced(pTree))
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

func (t *TreeNode) InOrder() {
	if t != nil {
		t.Left.InOrder()
		fmt.Print(t.Val, " ")
		t.Right.InOrder()
	}
}
