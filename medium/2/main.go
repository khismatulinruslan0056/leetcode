package main

import (
	"fmt"
	"time"
)

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored
in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the
sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example 1:


Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
Example 2:

Input: l1 = [0], l2 = [0]
Output: [0]
Example 3:

Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]


Constraints:

The number of nodes in each linked list is in the range [1, 100].
0 <= Node.val <= 9
It is guaranteed that the list represents a number that does not have leading zeros.*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sl1 := getFullNumbSlice(l1)
	sl2 := getFullNumbSlice(l2)

	result := sumSlices(sl1, sl2)

	return getSingleLinkedListFromNumb(result)
}

func sumList(l1 *ListNode, l2 *ListNode) *ListNode {
	resultL := &ListNode{Val: 0}
	current := resultL
	flag := false
	for {
		sum := current.Val
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		if sum > 9 {
			flag = true
			current.Val = sum % 10
		} else {
			flag = false
			current.Val = sum
		}

		if l1 == nil && l2 == nil && !flag {
			fmt.Println("pusto", l1, l2)
			break
		}
		current.Next = &ListNode{Val: 0}
		if flag {
			current.Next.Val++
		}
		current = current.Next

	}
	return resultL
}

func sumListOpt(l1 *ListNode, l2 *ListNode) *ListNode {
	resultL := &ListNode{}
	current := resultL
	carry := 0
	for {
		sum := current.Val
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		current.Val = sum % 10
		if l1 == nil && l2 == nil && carry < 1 {
			break
		}
		current.Next = &ListNode{Val: carry}
		current = current.Next
	}
	return resultL
}

func sumSlices(sl1, sl2 []int) []int {
	lens := len(sl1)
	if lens < len(sl2) {
		lens = len(sl2)
	}
	sumSlice := make([]int, lens)
	flag := false
	for i, _ := range sumSlice {
		sum := sumSlice[i]
		switch {
		case i < len(sl1) && i < len(sl2):
			sum += sl1[i] + sl2[i]
		case i >= len(sl1) && i < len(sl2):
			sum += sl2[i]
		case i < len(sl1) && i >= len(sl2):
			sum += sl1[i]
		}
		if sum > 9 {
			flag = true
			sumSlice[i] = sum % 10
		} else {
			flag = false
			sumSlice[i] = sum
		}
		if flag {
			if i+1 == len(sumSlice) {
				sumSlice = append(sumSlice, 1)
			} else {
				sumSlice[i+1]++
			}
		}

		fmt.Println(sumSlice[i], flag)
		fmt.Println("====================")
	}
	return sumSlice
}

func getSingleLinkedListFromNumb(numbSl []int) *ListNode {
	if len(numbSl) == 0 {
		return nil
	}
	head := &ListNode{Val: numbSl[0]}
	current := head
	for i := 1; i < len(numbSl); i++ {
		current.Next = &ListNode{Val: numbSl[i]}
		current = current.Next
	}
	return head
}

func getFullNumbSlice(l *ListNode) []int {
	sl1 := make([]int, 0)
	for l != nil {
		sl1 = append(sl1, l.Val)
		l = l.Next
	}
	return sl1
}

func reverseSingleLinkedList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

func print(l1 *ListNode) {
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
	//l1 := createList([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
	//l2 := createList([]int{5, 6, 4})
	//l1 := createList([]int{2, 4, 3})
	//l2 := createList([]int{5, 6, 4})
	l1 := createList([]int{9, 9, 9, 9, 9, 9, 9})
	l2 := createList([]int{9, 9, 9, 9})
	//					   8, 9, 9, 9, 0, 0, 0, 1
	list := sumListOpt(l1, l2)
	print(list)
	//print(reverseSingleLinkedList(list))

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

/*
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
Example 2:

Input: l1 = [0], l2 = [0]
Output: [0]
Example 3:

Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]*/
