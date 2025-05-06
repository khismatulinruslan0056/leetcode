package main

/*Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).



Example 1:


Input: root = [1,2,2,3,4,4,3]
Output: true
Example 2:


Input: root = [1,2,2,null,3,null,3]
Output: false*/

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left == nil {
		if root.Right == nil {
			return true
		}
		return false
	}

	if root.Right == nil {
		return false
	}
	currL := root
	currR := root
	stackL := []*TreeNode{currL}
	stackR := []*TreeNode{currR}
	for currR != nil || currL != nil || len(stackL) > 0 || len(stackR) > 0 {

		for currL != nil {
			stackL = append(stackL, currL)
			currL = currL.Left
		}
		for currR != nil {
			stackR = append(stackR, currR)
			currR = currR.Right
		}
		if len(stackR) != len(stackL) {
			return false
		}

		currR = stackR[len(stackR)-1]
		currL = stackL[len(stackL)-1]
		if currL.Val != currR.Val {
			return false
		}
		stackR = stackR[:len(stackR)-1]
		stackL = stackL[:len(stackL)-1]
		currL = currL.Right
		currR = currR.Left

	}
	return true
}

func main() {
	p := []interface{}{1, 2, 2, 3, 4, 4, 3}
	pTree := NewTree(p)
	fmt.Println(isSymmetric(pTree))

	p = []interface{}{1, 2, 2, nil, 3, nil, 3}
	pTree = NewTree(p)
	fmt.Println(isSymmetric(pTree))

	p = []interface{}{1, 2, 3}
	pTree = NewTree(p)
	fmt.Println(isSymmetric(pTree))
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
