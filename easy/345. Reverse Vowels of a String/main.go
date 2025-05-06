package main

import (
	"fmt"
	"slices"
)

/*
Given a string s, reverse only all the vowels in the string and return it.

The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.



Example 1:

Input: s = "IceCreAm"

Output: "AceCreIm"

Explanation:

The vowels in s are ['I', 'e', 'e', 'A']. On reversing the vowels, s becomes "AceCreIm".

Example 2:

Input: s = "leetcode"

Output: "leotcede"



Constraints:

1 <= s.length <= 3 * 105
s consist of printable ASCII characters.
Seen this question in a real interview before?
1/5

*/

func reverseVowels(s string) string {
	sByte := []byte(s)
	vowels := []byte("aeiouAEIOU")
	i, j := 0, len(sByte)-1
	for i < j {
		for i < j && !slices.Contains(vowels, sByte[i]) {
			i++
		}
		for i < j && !slices.Contains(vowels, sByte[j]) {
			j--
		}
		sByte[i], sByte[j] = sByte[j], sByte[i]
		i++
		j--
	}
	return string(sByte)
}

func main() {
	s := "IceCreAm"
	fmt.Println("Output: \"AceCreIm\"")
	fmt.Println(reverseVowels(s))
	s = "leetcode"
	fmt.Println("Output: \"leotcede\"")
	fmt.Println(reverseVowels(s))
}
