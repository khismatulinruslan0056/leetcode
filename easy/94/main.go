package main

/*
Given the root of a binary tree, return the inorder traversal of its nodes' values.



Example 1:

Input: root = [1,null,2,3]

Output: [1,3,2]

Explanation:



Example 2:

Input: root = [1,2,3,4,5,null,8,null,null,6,7,9]

Output: [4,2,6,5,7,1,3,9,8]

Explanation:



Example 3:

Input: root = []

Output: []

Example 4:

Input: root = [1]

Output: [1]



Constraints:

The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100


Follow up: Recursive solution is trivial, could you do it iteratively?*/

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val, Left: nil, Right: nil}
}

func (t *TreeNode) Insert(val int) *TreeNode {
	if t == nil {
		return &TreeNode{Val: val, Left: nil, Right: nil}
	}
	if val < t.Val {
		t.Left = t.Left.Insert(val)
	} else {
		t.Right = t.Right.Insert(val)
	}
	return t
}

func (t *TreeNode) Search(val int) bool {
	if t == nil {
		return false
	}
	if t.Val == val {
		return true
	}
	if val < t.Val {
		return t.Left.Search(val)
	} else {
		return t.Right.Search(val)
	}
}

func (t *TreeNode) InOrder() {
	if t != nil {
		t.Left.InOrder()
		fmt.Print(t.Val, " ")
		t.Right.InOrder()
	}
}

func (t *TreeNode) PreOrder() {
	if t != nil {
		fmt.Print(t.Val, " ")
		t.Left.InOrder()
		t.Right.InOrder()
	}
}

func (t *TreeNode) PostOrder() {
	if t != nil {
		t.Left.InOrder()
		fmt.Print(t.Val, " ")
		t.Right.InOrder()
	}
}

func main() {
	root := []interface{}{1, nil, 2, 3}
	rootTreeNode := NewTree(root)
	//rootTreeNode.InOrder2()
	fmt.Println(inorderTraversal(rootTreeNode))
	root = []interface{}{1, 2, 3, 4, 5, nil, 8, nil, nil, 6, 7, 9}
	rootTreeNode = NewTree(root)
	fmt.Println(inorderTraversal(rootTreeNode))
	root = []interface{}{}
	rootTreeNode = NewTree(root)
	fmt.Println(inorderTraversal(rootTreeNode))
	root = []interface{}{1}
	rootTreeNode = NewTree(root)
	fmt.Println(inorderTraversal(rootTreeNode))
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

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	curr := root
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, curr.Val)
		curr = curr.Right

	}
	return res
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
