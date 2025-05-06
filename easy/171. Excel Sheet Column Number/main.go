package main

import "fmt"

/*
Given a string columnTitle that represents the column title as appears in an Excel sheet, return its corresponding column number.

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

Input: columnTitle = "A"
Output: 1
Example 2:

Input: columnTitle = "AB"
Output: 28
Example 3:

Input: columnTitle = "ZY"
Output: 701


Constraints:

1 <= columnTitle.length <= 7
columnTitle consists only of uppercase English letters.
columnTitle is in the range ["A", "FXSHRXW"].
*/

func main() {
	columnTitle := "A"
	fmt.Println(1, columnTitle, titleToNumber(columnTitle))

	columnTitle = "AB"
	fmt.Println(28, columnTitle, titleToNumber(columnTitle))

	columnTitle = "ZY"
	fmt.Println(701, columnTitle, titleToNumber(columnTitle))
}

func titleToNumber(columnTitle string) int {
	result := 0
	for _, ch := range columnTitle {
		result = int(ch-'A'+1) + result*26
	}
	return result
}
