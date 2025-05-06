package main

import "fmt"

/*
Given a string s, find the length of the longest
substring
 without duplicate characters.



Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.


Constraints:

0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces.
*/

func lengthOfLongestSubstringWithArray(s string) int {
	arrayChInd := [128]int{}
	pos := 0
	maxLen := 0
	for i := 0; i < len(s); i++ {
		if arrayChInd[s[i]] > 0 && arrayChInd[s[i]] > pos {
			pos = arrayChInd[s[i]]
		}

		arrayChInd[s[i]] = i + 1
		maxLen = max(maxLen, i-pos+1)
	}
	return maxLen
}

func lengthOfLongestSubstringWithMap(s string) int {
	mapChInd := make(map[byte]int)
	maxLen := 0
	indexStart := 0
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	for i := 0; i < len(s); i++ {
		if _, ok := mapChInd[s[i]]; ok {
			if indexStart < mapChInd[s[i]]+1 {
				indexStart = mapChInd[s[i]] + 1
			}
		}

		mapChInd[s[i]] = i
		//str := s[indexStart : i+1]
		if maxLen < i+1-indexStart {
			maxLen = i + 1 - indexStart
		}
		//fmt.Println(str)
	}
	return maxLen
}

func main() {
	s := ""
	s = "abcabcbb"
	fmt.Println(s, 3, lengthOfLongestSubstringWithArray(s))
	fmt.Println("=================")
	s = "aab"
	fmt.Println(s, 2, lengthOfLongestSubstringWithArray(s))
	fmt.Println("=================")
	s = "pwwkew"
	fmt.Println(s, 3, lengthOfLongestSubstringWithArray(s))
	fmt.Println("=================")
	s = "dvdf"
	fmt.Println(s, 3, lengthOfLongestSubstringWithArray(s))
	fmt.Println("=================")
	s = "cdd"
	fmt.Println(s, 2, lengthOfLongestSubstringWithArray(s))
	fmt.Println("=================")
	s = "abba"
	fmt.Println(s, 2, lengthOfLongestSubstringWithArray(s))
	fmt.Println("=================")
	s = "ckilbkd"
	fmt.Println(s, 5, lengthOfLongestSubstringWithArray(s))
	fmt.Println("=================")
	s = "tmmzuxt"
	fmt.Println(s, 5, lengthOfLongestSubstringWithArray(s))
	fmt.Println("=================")
}
