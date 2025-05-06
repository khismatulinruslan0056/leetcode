package main

import (
	"fmt"
	"time"
)

/*
Given the head of a singly linked list, reverse the list, and return the reversed list.



Example 1:


Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]
Example 2:


Input: head = [1,2]
Output: [2,1]
Example 3:

Input: head = []
Output: []


Constraints:

The number of nodes in the list is the range [0, 5000].
-5000 <= Node.val <= 5000


Follow up: A linked list can be reversed either iteratively or recursively. Could you implement both?
*/

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
	list := []int{1, 2, 3, 4, 5}
	print(createList(list))
	result := reverseList(createList(list))
	print(result)

	list = []int{1, 2}
	print(createList(list))
	result = reverseList(createList(list))
	print(result)

	list = []int{}
	print(createList(list))
	result = reverseList(createList(list))
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

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	curr := head
	var prev *ListNode
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev

	//    if head == nil || head.Next == nil {
	//        return head
	//    }
	//    newHead := reverseList(head.Next)
	//    head.Next.Next = head  // Разворачиваем указатель
	//    head.Next = nil        // Убираем старую связь
	//    return newHead
}
