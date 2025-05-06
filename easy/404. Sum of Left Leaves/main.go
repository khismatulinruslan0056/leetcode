package main

/*
Given the root of a binary tree, return the sum of all left leaves.

A leaf is a node with no children. A left leaf is a leaf that is the left child of another node.



Example 1:


Input: root = [3,9,20,null,null,15,7]
Output: 24
Explanation: There are two left leaves in the binary tree, with values 9 and 15 respectively.
Example 2:

Input: root = [1]
Output: 0


Constraints:

The number of nodes in the tree is in the range [1, 1000].
-1000 <= Node.val <= 1000
Seen this question in a real interview before?
1/5 */

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		sum += root.Left.Val
	}
	sum += sumOfLeftLeaves(root.Left)
	sum += sumOfLeftLeaves(root.Right)
	return sum
}

func main() {
	p := []interface{}{-9, -3, 2, nil, 4, 4, 0, -6, nil, -5}
	pTree := NewTree(p)
	fmt.Println(sumOfLeftLeaves(pTree))
	p = []interface{}{1}
	pTree = NewTree(p)
	fmt.Println(sumOfLeftLeaves(pTree))

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
