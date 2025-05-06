package main

import (
	"fmt"
	"time"
)

/*
Given head, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.

Return true if there is a cycle in the linked list. Otherwise, return false.



Example 1:


Input: head = [3,2,0,-4], pos = 1
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).
Example 2:


Input: head = [1,2], pos = 0
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.
Example 3:


Input: head = [1], pos = -1
Output: false
Explanation: There is no cycle in the linked list.


Constraints:

The number of the nodes in the list is in the range [0, 104].
-105 <= Node.val <= 105
pos is -1 or a valid index in the linked-list.


Follow up: Can you solve it using O(1) (i.e. constant) memory?*/

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
	head := []int{3, 2, 0, -4}
	pos := 1
	list = createCircleList(head, pos)
	fmt.Println(hasCycle(list))
	head = []int{-21, 10, 17, 8, 4, 26, 5, 35, 33, -7, -16, 27, -12, 6, 29, -12, 5, 9, 20, 14, 14, 2, 13, -24, 21, 23, -21, 5}
	pos = -1
	list = createCircleList(head, pos)
	fmt.Println(hasCycle(list))
	head = []int{1}
	pos = -1
	list = createCircleList(head, pos)
	fmt.Println(hasCycle(list))

}

func createCircleList(values []int, position int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	current := head
	pointer := head
	for i, val := range values[1:] {
		current.Next = &ListNode{Val: val}
		if position == i+2 {
			pointer = current
		}
		current = current.Next
		if i == len(values[1:])-1 && position >= 0 && position < len(values) {
			current.Next = pointer
		}
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	//mapPointer := make(map[*ListNode]struct{})
	//current := head
	//for current != nil {
	//	if _, ok := mapPointer[current]; ok {
	//		return true
	//	}
	//	mapPointer[current] = struct{}{}
	//	current = current.Next
	//}
	//return false
	if head == nil || head.Next == nil {
		return false
	}
	fastPoint := head
	slowPoint := head
	for fastPoint != nil && fastPoint.Next != nil {

		fastPoint = fastPoint.Next.Next
		slowPoint = slowPoint.Next
		if fastPoint == slowPoint {
			return true
		}
	}
	return false
}
