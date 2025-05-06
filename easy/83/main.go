package main

import (
	"fmt"
	"time"
)

/*
Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.



Example 1:


Input: head = [1,1,2]
Output: [1,2]
Example 2:


Input: head = [1,1,2,3,3]
Output: [1,2,3]


Constraints:

The number of nodes in the list is in the range [0, 300].
-100 <= Node.val <= 100
The list is guaranteed to be sorted in ascending order.*/

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
	var list *ListNode
	//list := createList([]int{1, 1, 2})
	//print(list)
	//print(deleteDuplicates(list))
	list = createList([]int{1, 1, 2, 3, 3})
	print(list)
	print(deleteDuplicates(list))

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

func deleteDuplicates(head *ListNode) *ListNode {

	current := head

	for current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}
