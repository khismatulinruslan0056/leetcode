package main

import "fmt"

/*
 Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.



Example 1:

Input: s = "leetcode"

Output: 0

Explanation:

The character 'l' at index 0 is the first character that does not occur at any other index.

Example 2:

Input: s = "loveleetcode"

Output: 2

Example 3:

Input: s = "aabb"

Output: -1



Constraints:

1 <= s.length <= 105
s consists of only lowercase English letters.
Seen this question in a real interview before?
1/5*/

func firstUniqChar(s string) int {
	ascii := [128]int{}
	for i := 0; i < len(s); i++ {
		ascii[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if ascii[s[i]] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	s := "leetcode"
	fmt.Println(firstUniqChar(s))
	s = "loveleetcode"
	fmt.Println(firstUniqChar(s))

	s = "aabb"
	fmt.Println(firstUniqChar(s))
}
