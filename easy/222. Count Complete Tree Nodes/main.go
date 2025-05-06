package main

/*Given the root of a complete binary tree, return the number of the nodes in the tree.

According to Wikipedia, every level, except possibly the last, is completely filled in a complete binary tree, and all nodes in the last level are as far left as possible. It can have between 1 and 2h nodes inclusive at the last level h.

Design an algorithm that runs in less than O(n) time complexity.



Example 1:


Input: root = [1,2,3,4,5,6]
Output: 6
Example 2:

Input: root = []
Output: 0
Example 3:

Input: root = [1]
Output: 1


Constraints:

The number of nodes in the tree is in the range [0, 5 * 104].
0 <= Node.val <= 5 * 104
The tree is guaranteed to be complete.*/

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

func main() {
	p := []interface{}{1, 2, 3, 4, 5, 6}
	pTree := NewTree(p)
	fmt.Println(countNodes(pTree))

	p = []interface{}{}
	pTree = NewTree(p)
	fmt.Println(countNodes(pTree))

	p = []interface{}{1}
	pTree = NewTree(p)
	fmt.Println(countNodes(pTree))

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
