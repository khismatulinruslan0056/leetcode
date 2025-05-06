package main

import (
	"fmt"
)

/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.


Example 1:

Input: s = "()"

Output: true

Example 2:

Input: s = "()[]{}"

Output: true

Example 3:

Input: s = "(]"

Output: false

Example 4:

Input: s = "([])"

Output: true



Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.*/

var mapSkobok = map[byte]byte{
	'}': '{',
	']': '[',
	')': '(',
}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s)%2 == 1 {
		return false
	}
	stack := make([]byte, 0, len(s)/2)
	for i := 0; i < len(s); i++ {
		if s[i] == 40 || s[i] == 91 || s[i] == 123 {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != mapSkobok[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]

		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}

func main() {
	s := ""
	s = "()"
	fmt.Println(s, isValid(s))
	s = "()[]{}"
	fmt.Println(s, isValid(s))
	s = "(]"
	fmt.Println(s, isValid(s))
	s = "([])"
	fmt.Println(s, isValid(s))
	s = "(("
	fmt.Println(s, isValid(s))
	s = "(){}}{"
	fmt.Println(s, isValid(s))
	s = "({{{{}}}))"
	fmt.Println(s, isValid(s))
	s = "[([]])"
	fmt.Println(s, isValid(s))

}
