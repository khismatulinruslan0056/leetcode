package main

/*
 A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.



Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
Example 2:

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.
Example 3:

Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.


Constraints:

1 <= s.length <= 2 * 105
s consists only of printable ASCII characters.*/

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(s, isPalindrome(s))

	s = "race a car"
	fmt.Println(s, isPalindrome(s))

	s = "0P"
	fmt.Println(s, isPalindrome(s))
	s = ".,"
	fmt.Println(s, isPalindrome(s))
	s = "a,"
	fmt.Println(s, isPalindrome(s))
}

func isAlphanumeric(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !isAlphanumeric(rune(s[i])) {
			i++
		}
		for i < j && !isAlphanumeric(rune(s[j])) {
			j--
		}

		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}
