package main

import "fmt"

/*
 The Fibonacci numbers, commonly denoted F(n) form a sequence, called the Fibonacci sequence, such that each number is the sum of the two preceding ones, starting from 0 and 1. That is,

F(0) = 0, F(1) = 1
F(n) = F(n - 1) + F(n - 2), for n > 1.
Given n, calculate F(n).



Example 1:

Input: n = 2
Output: 1
Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.
Example 2:

Input: n = 3
Output: 2
Explanation: F(3) = F(2) + F(1) = 1 + 1 = 2.
Example 3:

Input: n = 4
Output: 3
Explanation: F(4) = F(3) + F(2) = 2 + 1 = 3.


Constraints:

0 <= n <= 30*/

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	fib1, fib2 := 0, 1
	//return fib(n-1) + fib(n-2)
	for i := 2; i <= n; i++ {
		fib1, fib2 = fib2, fib1+fib2
	}
	return fib2
}

func main() {
	n := 2
	fmt.Println(fib(n))
	n = 3
	fmt.Println(fib(n))
	n = 4
	fmt.Println(fib(n))
}
