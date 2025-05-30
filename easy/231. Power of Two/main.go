package main

import "fmt"

/*
Given an integer n, return true if it is a power of two. Otherwise, return false.

An integer n is a power of two, if there exists an integer x such that n == 2x.



Example 1:

Input: n = 1
Output: true
Explanation: 20 = 1
Example 2:

Input: n = 16
Output: true
Explanation: 24 = 16
Example 3:

Input: n = 3
Output: false


Constraints:

-231 <= n <= 231 - 1


Follow up: Could you solve it without loops/recursion?*/

func isPowerOfTwo(n int) bool {
	//if n%2 == 1 && n != 1 || n <= 0 {
	//	return false
	//}
	//if n <= 2 {
	//	return true
	//}
	//return isPowerOfTwo(n / 2)

	if n <= 0 {
		return false
	}
	if n <= 2 {
		return true
	}
	for n > 2 {
		if n%2 == 1 {
			return false
		}
		n = n / 2
	}
	return true
}

func main() {
	n := 1
	fmt.Println(isPowerOfTwo(n))
	n = 16
	fmt.Println(isPowerOfTwo(n))
	n = 3
	fmt.Println(isPowerOfTwo(n))
}
