package main

import (
	"fmt"
	"strings"
)

/*
 Given a pattern and a string s, find if s follows the same pattern.

Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s. Specifically:

Each letter in pattern maps to exactly one unique word in s.
Each unique word in s maps to exactly one letter in pattern.
No two letters map to the same word, and no two words map to the same letter.


Example 1:

Input: pattern = "abba", s = "dog cat cat dog"

Output: true

Explanation:

The bijection can be established as:

'a' maps to "dog".
'b' maps to "cat".
Example 2:

Input: pattern = "abba", s = "dog cat cat fish"

Output: false

Example 3:

Input: pattern = "aaaa", s = "dog cat cat dog"

Output: false



Constraints:

1 <= pattern.length <= 300
pattern contains only lower-case English letters.
1 <= s.length <= 3000
s contains only lowercase English letters and spaces ' '.
s does not contain any leading or trailing spaces.
All the words in s are separated by a single space.*/

func wordPattern(pattern string, s string) bool {
	sSl := strings.Split(s, " ")
	if len(pattern) != len(sSl) {
		return false
	}
	mapPattern := make(map[byte]string)
	mapS := make(map[string]byte)
	for i, str := range sSl {
		if v, ok := mapS[str]; ok && v != pattern[i] {
			return false
		}
		if v, ok := mapPattern[pattern[i]]; ok && v != sSl[i] {
			return false
		}
		mapPattern[pattern[i]] = str
		mapS[str] = pattern[i]
	}
	return true
}

func main() {
	pattern := "abba"
	s := "dog cat cat dog"
	fmt.Println(wordPattern(pattern, s))
	pattern = "abba"
	s = "dog cat cat fish"
	fmt.Println(wordPattern(pattern, s))
	pattern = "aaaa"
	s = "dog cat cat dog"
	fmt.Println(wordPattern(pattern, s))
}
