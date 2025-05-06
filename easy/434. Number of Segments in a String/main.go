package main

import (
	"fmt"
	"unicode"
)

/*
 Given a string s, return the number of segments in the string.

A segment is defined to be a contiguous sequence of non-space characters.



Example 1:

Input: s = "Hello, my name is John"
Output: 5
Explanation: The five segments are ["Hello,", "my", "name", "is", "John"]
Example 2:

Input: s = "Hello"
Output: 1


Constraints:

0 <= s.length <= 300
s consists of lowercase and uppercase English letters, digits, or one of the following characters "!@#$%^&*()_+-=',.:".
The only space character in s is ' '.*/

func countSegments(s string) int {
	if s == "" {
		return 0
	}
	count := 0
	countCH := 0
	for _, ch := range s {
		if unicode.IsSpace(ch) {
			if countCH > 0 {
				count++
				countCH = 0
				continue
			}
			continue
		}
		countCH++
	}
	if countCH > 0 {
		count++
	}
	return count
}

func main() {
	s := "Hello, my name is John"
	fmt.Println(countSegments(s))
	s = "Hello"
	fmt.Println(countSegments(s))
}
