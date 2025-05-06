package main

import (
	"fmt"
	"strings"
)

/*
 Given an array of strings words, return the words that can be typed using letters of the alphabet on only one row of American keyboard like the image below.

Note that the strings are case-insensitive, both lowercased and uppercased of the same letter are treated as if they are at the same row.

In the American keyboard:

the first row consists of the characters "qwertyuiop",
the second row consists of the characters "asdfghjkl", and
the third row consists of the characters "zxcvbnm".



Example 1:

Input: words = ["Hello","Alaska","Dad","Peace"]

Output: ["Alaska","Dad"]

Explanation:

Both "a" and "A" are in the 2nd row of the American keyboard due to case insensitivity.

Example 2:

Input: words = ["omk"]

Output: []

Example 3:

Input: words = ["adsdf","sfd"]

Output: ["adsdf","sfd"]



Constraints:

1 <= words.length <= 20
1 <= words[i].length <= 100
words[i] consists of English letters (both lowercase and uppercase). */

func findWords(words []string) []string {
	rows := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
	result := make([]string, 0)
	for _, row := range rows {
		for _, word := range words {
			if exist(word, row) {
				result = append(result, word)
			}
		}
	}
	return result
}

func exist(word string, row string) bool {
	word = strings.ToLower(word)
	for i := 0; i < len(word); i++ {
		if !strings.Contains(row, string(word[i])) {
			return false
		}
	}
	return true
}

func main() {
	words := []string{"Hello", "Alaska", "Dad", "Peace"}
	fmt.Println(findWords(words))
	words = []string{"omk"}
	fmt.Println(findWords(words))
	words = []string{"adsdf", "sfd"}
	fmt.Println(findWords(words))
}
