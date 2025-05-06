package main

import (
	"fmt"
)

/*
 Given an integer num, return a string of its base 7 representation.



Example 1:

Input: num = 100
Output: "202"
Example 2:

Input: num = -7
Output: "-10"


Constraints:

-107 <= num <= 107
Seen this question in a real interview before?
1/5*/

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	res := make([]byte, 0)
	negative := num < 0
	if negative {
		num *= -1
	}
	for num != 0 {
		res = append([]byte{byte('0' + num%7)}, res...)
		num /= 7
	}

	if negative {
		res = append([]byte{'-'}, res...)
	}
	return string(res)
}

func main() {
	num := 100
	fmt.Println(convertToBase7(num))
	num = -7
	fmt.Println(convertToBase7(num))
}
