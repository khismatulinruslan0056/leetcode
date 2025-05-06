package main

import (
	"fmt"
)

/*
Given a 32-bit integer num, return a string representing its hexadecimal representation. For negative integers, two’s complement method is used.

All the letters in the answer string should be lowercase characters, and there should not be any leading zeros in the answer except for the zero itself.

Note: You are not allowed to use any built-in library method to directly solve this problem.



Example 1:

Input: num = 26
Output: "1a"
Example 2:

Input: num = -1
Output: "ffffffff"


Constraints:

-231 <= num <= 231 - 1
Seen this question in a real interview before?
1/5
*/

func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	unsignedNum := uint32(num)

	hexChars := []byte("0123456789abcdef")
	res := make([]byte, 0, 8)

	for unsignedNum > 0 {
		digit := unsignedNum & 0xf
		res = append([]byte{hexChars[digit]}, res...)
		unsignedNum >>= 4
	}
	return string(res)
	//res := ""
	//map16 := map[int]string{
	//	10: "A",
	//	11: "B",
	//	12: "C",
	//	13: "D",
	//	14: "E",
	//	15: "F",
	//}
	//if num < 0 {
	//	num *= -1
	//	result := ""
	//	for num > 0 {
	//		result = strconv.Itoa(num%2) + result
	//		num /= 2
	//	}
	//	for len(result) != 8 {
	//		result = "0" + result
	//	}
	//	result2 := ""
	//	for i := 0; i < len(result); i++ {
	//		if result[i] == '0' {
	//			result2 += "1"
	//		} else {
	//			result2 += "0"
	//		}
	//	}
	//	curr, _ := strconv.ParseInt(result2, 2, 32)
	//	num = int(curr + 1)
	//
	//}
	//for num != 0 {
	//	a := num % 16
	//	if a < 10 {
	//		res = strconv.Itoa(a) + res
	//	} else {
	//		res = map16[a] + res
	//	}
	//	num /= 16
	//}
	//
	//return res
}

func main() {
	num := 26
	fmt.Println(toHex(num))
	num = -1
	fmt.Println(toHex(num))
}
