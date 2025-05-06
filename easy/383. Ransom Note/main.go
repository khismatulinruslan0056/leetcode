package main

import "fmt"

/*Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.

Each letter in magazine can only be used once in ransomNote.



Example 1:

Input: ransomNote = "a", magazine = "b"
Output: false
Example 2:

Input: ransomNote = "aa", magazine = "ab"
Output: false
Example 3:

Input: ransomNote = "aa", magazine = "aab"
Output: true


Constraints:

1 <= ransomNote.length, magazine.length <= 105
ransomNote and magazine consist of lowercase English letters.
Seen this question in a real interview before?
1/5*/

func canConstruct(ransomNote string, magazine string) bool {
	ascii := [128]int{}
	for i := 0; i < len(magazine); i++ {
		ascii[magazine[i]]++
	}
	for i := 0; i < len(ransomNote); i++ {
		if ascii[ransomNote[i]] == 0 {
			return false
		}
		ascii[ransomNote[i]]--
	}
	return true
}

func main() {
	ransomNote := "a"
	magazine := "b"
	fmt.Println(canConstruct(ransomNote, magazine))
	ransomNote = "aa"
	magazine = "ab"
	fmt.Println(canConstruct(ransomNote, magazine))
	ransomNote = "aa"
	magazine = "aab"
	fmt.Println(canConstruct(ransomNote, magazine))
}
