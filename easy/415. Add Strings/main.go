package main

import (
	"fmt"
	"strconv"
)

/*
 Given two non-negative integers, num1 and num2 represented as string, return the sum of num1 and num2 as a string.

You must solve the problem without using any built-in library for handling large integers (such as BigInteger). You must also not convert the inputs to integers directly.



Example 1:

Input: num1 = "11", num2 = "123"
Output: "134"
Example 2:

Input: num1 = "456", num2 = "77"
Output: "533"
Example 3:

Input: num1 = "0", num2 = "0"
Output: "0"


Constraints:

1 <= num1.length, num2.length <= 104
num1 and num2 consist of only digits.
num1 and num2 don't have any leading zeros except for the zero itself.*/

func addStrings(num1 string, num2 string) string {
	res := ""
	carry := 0
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || carry > 0; i, j = i-1, j-1 {
		n := carry
		if i >= 0 {
			a, _ := strconv.Atoi(string(num1[i]))
			n += a
		}
		if j >= 0 {
			b, _ := strconv.Atoi(string(num2[j]))
			n += b
		}

		carry = n / 10
		res = strconv.Itoa(n%10) + res

	}
	return res
}

func main() {
	num1 := "123456789"
	num2 := "987654321"
	fmt.Println(addStrings(num1, num2))

	num1 = "456"
	num2 = "877"
	fmt.Println(addStrings(num1, num2))

	num1 = "0"
	num2 = "0"
	fmt.Println(addStrings(num1, num2))
}
