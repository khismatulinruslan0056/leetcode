package main

/*
 Given the root of a binary search tree (BST) with duplicates, return all the mode(s) (i.e., the most frequently occurred element) in it.

If the tree has more than one mode, return them in any order.

Assume a BST is defined as follows:

The left subtree of a node contains only nodes with keys less than or equal to the node's key.
The right subtree of a node contains only nodes with keys greater than or equal to the node's key.
Both the left and right subtrees must also be binary search trees.


Example 1:


Input: root = [1,null,2,2]
Output: [2]
Example 2:

Input: root = [0]
Output: [0]


Constraints:

The number of nodes in the tree is in the range [1, 104].
-105 <= Node.val <= 105


Follow up: Could you do that without using any extra space? (Assume that the implicit stack space
incurred due to recursion does not count).*/

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var (
		maxCounts = 0
		count     = 0
		curr      = 0
		modes     = make([]int, 0)
	)

	updateMode := func(val int) {
		if val != curr {
			curr = val
			count = 0
		}
		count++

		if count > maxCounts {
			maxCounts = count
			modes = []int{val}
		} else if count == maxCounts {
			modes = append(modes, val)
		}
	}
	var inOrder func(*TreeNode)
	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inOrder(root.Left)
		updateMode(root.Val)
		inOrder(root.Right)
	}
	inOrder(root)
	return modes
}

func main() {
	p := []interface{}{1, nil, 2, 2}
	pTree := NewTree(p)
	fmt.Println(findMode(pTree))

	p = []interface{}{0}
	pTree = NewTree(p)
	fmt.Println(findMode(pTree))

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
