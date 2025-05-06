package main

import (
	"fmt"
	"time"
)

/*
 Given the head of a linked list and an integer val, remove all the nodes of the linked list that has Node.val == val, and return the new head.



Example 1:


Input: head = [1,2,6,3,4,5,6], val = 6
Output: [1,2,3,4,5]
Example 2:

Input: head = [], val = 1
Output: []
Example 3:

Input: head = [7,7,7,7], val = 7
Output: []


Constraints:

The number of nodes in the list is in the range [0, 104].
1 <= Node.val <= 50
0 <= val <= 50*/

func print(l1 *ListNode) {
	if l1 == nil {
		fmt.Println(nil)
		return
	}
	fmt.Print(l1.Val)
	if l1.Next != nil {
		fmt.Print("->")
	}
	time.Sleep(time.Millisecond)
	if l1.Next == nil {
		fmt.Println()
		return
	}
	print(l1.Next)
}

func main() {
	list := []int{1, 2, 6, 3, 4, 5, 6}
	n := 6
	result := removeElements(createList(list), n)
	print(result)

	list = []int{1, 2, 2, 1}
	n = 2
	result = removeElements(createList(list), n)
	print(result)

	list = []int{7, 7, 7, 7}
	n = 7
	result = removeElements(createList(list), n)
	print(result)

}

func createList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	current := head
	for _, val := range values[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {

	//1, 2, 6, 3, 4, 5, 6
	for head != nil && head.Val == val {
		head = head.Next
	}
	if head == nil {
		return nil
	}
	curr := head
	prev := curr
	for curr != nil {
		if curr.Val == val {
			prev.Next = curr.Next

		} else {
			prev = curr
		}
		curr = curr.Next
	}
	return head
}
