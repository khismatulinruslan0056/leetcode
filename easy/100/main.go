package main

/*Given the roots of two binary trees p and q, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.



Example 1:


Input: p = [1,2,3], q = [1,2,3]
Output: true
Example 2:


Input: p = [1,2], q = [1,null,2]
Output: false
Example 3:


Input: p = [1,2,1], q = [1,1,2]
Output: false


Constraints:

The number of nodes in both trees is in the range [0, 100].
-104 <= Node.val <= 104*/

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil {
		if q == nil {
			return true
		}
		return false
	}
	if q == nil {
		return false
	}
	stackP := make([]*TreeNode, 0)
	stackQ := make([]*TreeNode, 0)
	currP := p
	currQ := q
	for currP != nil || len(stackP) > 0 || currQ != nil || len(stackQ) > 0 {
		for currP != nil {
			stackP = append(stackP, currP)
			currP = currP.Left
		}
		for currQ != nil {
			stackQ = append(stackQ, currQ)
			currQ = currQ.Left
		}
		if len(stackP) != len(stackQ) {
			return false
		}
		currP = stackP[len(stackP)-1]
		currQ = stackQ[len(stackQ)-1]
		if currP.Val != currQ.Val {
			return false
		}
		stackP = stackP[:len(stackP)-1]
		stackQ = stackQ[:len(stackQ)-1]
		currP = currP.Right
		currQ = currQ.Right
	}
	return true
}

func main() {
	p := []interface{}{1, 2, 3}
	q := []interface{}{1, 2, 3}
	pTree := NewTree(p)
	qTree := NewTree(q)
	fmt.Println(isSameTree(pTree, qTree))
	p = []interface{}{1, 2}
	q = []interface{}{1, nil, 2}
	pTree = NewTree(p)
	qTree = NewTree(q)
	fmt.Println(isSameTree(pTree, qTree))

	p = []interface{}{1, 2, 1}
	q = []interface{}{1, 1, 2}
	pTree = NewTree(p)
	qTree = NewTree(q)
	fmt.Println(isSameTree(pTree, qTree))
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
