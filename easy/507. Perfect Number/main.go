package main

import "fmt"

/*
A perfect number is a positive integer that is equal to the sum of its positive divisors, excluding the number itself. A divisor of an integer x is an integer that can divide x evenly.

Given an integer n, return true if n is a perfect number, otherwise return false.



Example 1:

Input: num = 28
Output: true
Explanation: 28 = 1 + 2 + 4 + 7 + 14
1, 2, 4, 7, and 14 are all divisors of 28.
Example 2:

Input: num = 7
Output: false


Constraints:*/

func checkPerfectNumber(num int) bool {
	//if num == 1 || num == 2 {
	//	return false
	//}
	//res := num
	//for i := num / 2; i > 0; i-- {
	//	if num%i == 0 {
	//		res -= i
	//	}
	//}
	//
	//return res == 0

	if num == 1 {
		return false
	}
	sum := 1

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			sum += i
			if num/i != i {
				sum += num / i
			}
		}
	}
	return sum == num
}

func main() {
	num := 28
	fmt.Println(checkPerfectNumber(num))
	num = 7
	fmt.Println(checkPerfectNumber(num))
}
