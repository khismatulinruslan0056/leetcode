package main

import "fmt"

/*
	Given a string s which consists of lowercase or uppercase letters, return the length of the longest palindrome that can be built with those letters.

Letters are case sensitive, for example, "Aa" is not considered a palindrome.

Example 1:

Input: s = "abccccdd"
Output: 7
Explanation: One longest palindrome that can be built is "dccaccd", whose length is 7.
Example 2:

Input: s = "a"
Output: 1
Explanation: The longest palindrome that can be built is "a", whose length is 1.

Constraints:

1 <= s.length <= 2000
s consists of lowercase and/or uppercase English letters only.
*/
func longestPalindrome(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	ascii := [128]int{}
	for i := 0; i < len(s); i++ {
		ascii[s[i]]++
	}
	longest := 0
	for _, b := range ascii {
		if b > 1 {
			longest += b / 2 * 2
		}
	}
	if longest < len(s) {
		longest++
	}
	return longest
}

func main() {
	s := "abccccdd"
	fmt.Println(longestPalindrome(s))
	s = "ccc"
	fmt.Println(longestPalindrome(s))
}
