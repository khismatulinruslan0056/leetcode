package main

import (
	"fmt"
	"strings"
)

/*We define the usage of capitals in a word to be right when one of the following cases holds:

All letters in this word are capitals, like "USA".
All letters in this word are not capitals, like "leetcode".
Only the first letter in this word is capital, like "Google".
Given a string word, return true if the usage of capitals in it is right.



Example 1:

Input: word = "USA"
Output: true
Example 2:

Input: word = "FlaG"
Output: false


Constraints:

1 <= word.length <= 100
word consists of lowercase and uppercase English letters.*/

func detectCapitalUse(word string) bool {
	return strings.ToLower(word) == word || strings.ToUpper(word) == word || strings.ToLower(word[1:]) == word[1:]
}

func main() {
	word := "USA"
	fmt.Println(detectCapitalUse(word))
	word = "FlaG"
	fmt.Println(detectCapitalUse(word))
}
