package main

import "fmt"

/*
 Given two strings s and t, return true if t is an anagram of s, and false otherwise.



Example 1:

Input: s = "anagram", t = "nagaram"

Output: true

Example 2:

Input: s = "rat", t = "car"

Output: false



Constraints:

1 <= s.length, t.length <= 5 * 104
s and t consist of lowercase English letters.


Follow up: What if the inputs contain Unicode characters? How would you adapt your solution to such a case?*/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	ascii := [128]int{}
	for i := 0; i < len(s); i++ {
		ascii[s[i]]++
		ascii[t[i]]--
	}
	for _, num := range ascii {
		if num != 0 {
			return false
		}
	}
	return true
}

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))
	s = "rat"
	t = "car"
	fmt.Println(isAnagram(s, t))
}
