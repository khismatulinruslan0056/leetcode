package main

import "fmt"

/*
 Given an integer columnNumber, return its corresponding column title as it appears in an Excel sheet.

For example:

A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28
...


Example 1:

Input: columnNumber = 1
Output: "A"
Example 2:

Input: columnNumber = 28
Output: "AB"
Example 3:

Input: columnNumber = 701
Output: "ZY"


Constraints:

1 <= columnNumber <= 231 - 1*/

func main() {
	columnNumber := 1
	fmt.Println(columnNumber, "A", convertToTitle(columnNumber))
	columnNumber = 28
	fmt.Println(columnNumber, "AB", convertToTitle(columnNumber))
	columnNumber = 701
	fmt.Println(columnNumber, "ZY", convertToTitle(columnNumber))
	columnNumber = 52
	fmt.Println(columnNumber, "ZY", convertToTitle(columnNumber))
}

func convertToTitle(columnNumber int) string {
	//init := "ZABCDEFGHIJKLMNOPQRSTUVWXY"
	//result := make([]byte, 0)
	//
	//for columnNumber > 0 {
	//	result = append([]byte{init[columnNumber%26]}, result...)
	//	if columnNumber%26 == 0 {
	//		columnNumber--
	//	}
	//	columnNumber /= 26
	//
	//}
	//return string(result)
	result := ""
	for columnNumber > 0 {
		columnNumber--
		remainder := columnNumber % 26
		result = string('A'+remainder) + result
		columnNumber /= 26
	}
	return result
}
