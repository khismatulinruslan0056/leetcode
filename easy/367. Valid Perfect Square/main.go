package main

import "fmt"

/*
 Given a positive integer num, return true if num is a perfect square or false otherwise.

A perfect square is an integer that is the square of an integer. In other words, it is the product of some integer with itself.

You must not use any built-in library function, such as sqrt.



Example 1:

Input: num = 16
Output: true
Explanation: We return true because 4 * 4 = 16 and 4 is an integer.
Example 2:

Input: num = 14
Output: false
Explanation: We return false because 3.742 * 3.742 = 14 and 3.742 is not an integer.


Constraints:

1 <= num <= 231 - 1
Seen this question in a real interview before?
1/5*/

func isPerfectSquare(num int) bool {
	//if num == 1 {
	//	return true
	//}
	//for i := 2; i*i <= num; i++ {
	//	if i*i == num {
	//		return true
	//	}
	//}
	//return false
	if num < 1 {
		return false
	}

	start, end := 1, num
	for start <= end {
		mid := (start + end) / 2
		square := mid * mid
		if square == num {
			return true
		} else if square < num {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return false
}

func main() {
	num := 16
	fmt.Println(isPerfectSquare(num))
	num = 14
	fmt.Println(isPerfectSquare(num))
}
