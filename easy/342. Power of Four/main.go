package main

import "fmt"

/*
 Given an integer n, return true if it is a power of three. Otherwise, return false.

An integer n is a power of three, if there exists an integer x such that n == 3x.



Example 1:

Input: n = 27
Output: true
Explanation: 27 = 33
Example 2:

Input: n = 0
Output: false
Explanation: There is no x where 3x = 0.
Example 3:

Input: n = -1
Output: false
Explanation: There is no x where 3x = (-1).


Constraints:

-231 <= n <= 231 - 1


Follow up: Could you solve it without loops/recursion?*/

func isPowerOfFour(n int) bool {
	if n < 1 {
		return false
	}
	for n >= 4 {
		if n%4 != 0 {
			return false
		}
		n /= 4
	}

	return n == 1
}

func main() {
	n := 16
	fmt.Println(isPowerOfFour(n))
	n = 5
	fmt.Println(isPowerOfFour(n))
	n = 1
	fmt.Println(isPowerOfFour(n))
}
