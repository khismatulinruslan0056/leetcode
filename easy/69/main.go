package main

import (
	"fmt"
)

/*
Given a non-negative integer x, return the square root of x rounded down to the nearest integer. The returned integer should be non-negative as well.

You must not use any built-in exponent function or operator.

For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.


Example 1:

Input: x = 4
Output: 2
Explanation: The square root of 4 is 2, so we return 2.
Example 2:

Input: x = 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since we round it down to the nearest integer, 2 is returned.


Constraints:

0 <= x <= 231 - 1*/

func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	left := 0
	right := x

	for left <= right {
		mid := (left + right) / 2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right

}

func main() {
	x := 4
	fmt.Println("Output: 2")
	fmt.Println(mySqrt(x))
	x = 8
	fmt.Println("Output: 2")
	fmt.Println(mySqrt(x))
}
