package main

import "fmt"

/*
 Given an integer num, repeatedly add all its digits until the result has only one digit, and return it.



Example 1:

Input: num = 38
Output: 2
Explanation: The process is
38 --> 3 + 8 --> 11
11 --> 1 + 1 --> 2
Since 2 has only one digit, return it.
Example 2:

Input: num = 0
Output: 0


Constraints:

0 <= num <= 231 - 1


Follow up: Could you do it without any loop/recursion in O(1) runtime?*/

func addDigits(num int) int {
	//if num < 10 {
	//	return num
	//}
	//sum := 0
	//for {
	//	sum += num % 10
	//	num /= 10
	//	if num == 0 {
	//		num = sum
	//		if num < 10 {
	//			return num
	//		}
	//		sum = 0
	//	}
	//}
	//return sum
	if num < 10 {
		return num
	}
	if num%9 == 0 {
		return 9
	}
	return num % 9
}

func main() {
	num := 38
	fmt.Println(num, addDigits(num))
	num = 0
	fmt.Println(num, addDigits(num))

}
