package main

import (
	"fmt"
)

/*You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer. The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading 0's.

Increment the large integer by one and return the resulting array of digits.



Example 1:

Input: digits = [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Incrementing by one gives 123 + 1 = 124.
Thus, the result should be [1,2,4].
Example 2:

Input: digits = [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
Incrementing by one gives 4321 + 1 = 4322.
Thus, the result should be [4,3,2,2].
Example 3:

Input: digits = [9]
Output: [1,0]
Explanation: The array represents the integer 9.
Incrementing by one gives 9 + 1 = 10.
Thus, the result should be [1,0].*/

func plusOne(digits []int) []int {
	//res := make([]int, len(digits))
	//res = digits
	//res[len(digits)-1]++
	//if digits[len(digits)-1] != 10 {
	//	return res
	//}
	//for i := len(res) - 1; i >= 0; i-- {
	//	if res[i] > 9 {
	//		res[i] = 0
	//		if i == 0 {
	//			res = append([]int{1}, res...)
	//		} else {
	//			res[i-1]++
	//		}
	//	}
	//}
	//return res
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			return digits
		}
	}

	return append([]int{1}, digits...)
}

func main() {
	digits := []int{1, 2, 3}
	fmt.Println(digits)
	fmt.Println("Output: [1,2,4]")
	fmt.Println(plusOne(digits))

	digits = []int{4, 3, 2, 1}
	fmt.Println(digits)
	fmt.Println("Output: [4,3,2,2]")
	fmt.Println(plusOne(digits))
	digits = []int{9}
	fmt.Println(digits)
	fmt.Println("Output: [1,0]")
	fmt.Println(plusOne(digits))

	digits = []int{9, 8, 9}
	fmt.Println(digits)
	fmt.Println("Output: [9,9,0]")
	fmt.Println(plusOne(digits))

	a := '1'
	fmt.Println((a) % 49)
	fmt.Println((a + a) % 49)
	fmt.Println((a + a + a) % 49)

}
