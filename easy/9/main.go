package main

import "fmt"

/*
9. Palindrome Number
Easy

Topics
Companies

Hint
Given an integer x, return true if x is a
palindrome
, and false otherwise.



Example 1:

Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.
Example 2:

Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.


Constraints:

-231 <= x <= 231 - 1


Follow up: Could you solve it without converting the integer to a string?*/

func isPalindrome(x int) bool {
	//if x < 0 {
	//	return false
	//}
	//if x < 10 {
	//	return true
	//}
	//sl := make([]int, 0)
	//for x != 0 {
	//	sl = append(sl, x%10)
	//	x /= 10
	//}
	//for i := 0; i < len(sl)/2+1; i++ {
	//	if sl[i] != sl[len(sl)-1-i] {
	//		return false
	//	}
	//}
	//return true

	if x < 0 || x%10 == 0 {
		return false
	}
	if x < 10 {
		return true
	}

	reverse := 0
	original := x
	for original > reverse {
		reverse = reverse*10 + original%10
		original /= 10
	}
	if original == reverse || reverse/10 == original {
		return true
	}
	return false
}

func main() {
	x := 0
	x = 121
	fmt.Println(x, isPalindrome(x))
	x = -121
	fmt.Println(x, isPalindrome(x))
	x = 10
	fmt.Println(x, isPalindrome(x))
}
