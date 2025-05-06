package main

import "fmt"

/*
 Given a string s, check if it can be constructed by taking a substring of it and appending multiple copies of the substring together.



Example 1:

Input: s = "abab"
Output: true
Explanation: It is the substring "ab" twice.
Example 2:

Input: s = "aba"
Output: false
Example 3:

Input: s = "abcabcabcabc"
Output: true
Explanation: It is the substring "abc" four times or the substring "abcabc" twice.


Constraints:

1 <= s.length <= 104
s consists of lowercase English letters.*/

func repeatedSubstringPattern(s string) bool {
	double := s + s
	for i := 1; i <= len(s)/2; i++ {
		a := double[i : i+len(s)]
		_ = a
		if double[i:i+len(s)] == s {
			return true
		}
	}

	return false
}

func main() {
	s := "abac"
	fmt.Println(repeatedSubstringPattern(s))

	s = "aba"
	fmt.Println(repeatedSubstringPattern(s))

	s = "abcabcabcabc"
	fmt.Println(repeatedSubstringPattern(s))
}
