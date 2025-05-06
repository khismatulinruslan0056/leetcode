package main

import (
	"fmt"
	"time"
)

/*
Given the head of a singly linked list, return true if it is a palindrome or false otherwise.



Example 1:


Input: head = [1,2,2,1]
Output: true
Example 2:


Input: head = [1,2]
Output: false


Constraints:

The number of nodes in the list is in the range [1, 105].
0 <= Node.val <= 9

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
	list := []int{1, 1, 2, 1}
	print(createList(list))
	fmt.Println(isPalindrome(createList(list)))

	list = []int{1, 2, 2, 1}
	print(createList(list))
	fmt.Println(isPalindrome(createList(list)))

	list = []int{1, 2}
	print(createList(list))
	fmt.Println(isPalindrome(createList(list)))
	//1,1,2,1

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

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	curr := slow
	var prev *ListNode

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	first, second := head, prev
	for second != nil {
		if first.Val != second.Val {
			return false
		}
		first = first.Next
		second = second.Next
	}

	return true
}
