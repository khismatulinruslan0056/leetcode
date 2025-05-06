package main

/*Given the root of a binary tree, return its maximum depth.

A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.



Example 1:


Input: root = [3,9,20,null,null,15,7]
Output: 3
Example 2:

Input: root = [1,null,2]
Output: 2


Constraints:

The number of nodes in the tree is in the range [0, 104].
-100 <= Node.val <= 100*/

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepthRec(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	return 1 + max(leftDepth, rightDepth)
}

func maxDepth(root *TreeNode) int {
	queue := []*TreeNode{root}
	depth := 0
	for len(queue) > 0 {
		lengthCurr := len(queue)
		depth++
		for i := 0; i < lengthCurr; i++ {
			curr := queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
	}
	return depth
}

func main() {
	p := []interface{}{3, 9, 20, nil, nil, 15, 7}
	pTree := NewTree(p)
	fmt.Println(maxDepth(pTree))

	p = []interface{}{1, nil, 2}
	pTree = NewTree(p)
	fmt.Println(maxDepth(pTree))
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
