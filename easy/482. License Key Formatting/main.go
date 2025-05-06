package main

import (
	"fmt"
)

/*
 You are given a license key represented as a string s that consists of only alphanumeric characters and dashes. The string is separated into n + 1 groups by n dashes. You are also given an integer k.

We want to reformat the string s such that each group contains exactly k characters, except for the first group, which could be shorter than k but still must contain at least one character. Furthermore, there must be a dash inserted between two groups, and you should convert all lowercase letters to uppercase.

Return the reformatted license key.



Example 1:

Input: s = "5F3Z-2e-9-w", k = 4
Output: "5F3Z-2E9W"
Explanation: The string s has been split into two parts, each part has 4 characters.
Note that the two extra dashes are not needed and can be removed.
Example 2:

Input: s = "2-5g-3-J", k = 2
Output: "2-5G-3J"
Explanation: The string s has been split into three parts, each part has 2 characters except the first part as it could be shorter as mentioned above.


Constraints:

1 <= s.length <= 105
s consists of English letters, digits, and dashes '-'.
1 <= k <= 104*/

func licenseKeyFormatting(s string, k int) string {

	n := 0
	for _, ch := range s {
		if ch != '-' {
			n++
		}
	}
	if n == 0 {
		return ""
	}
	fullGroups := n / k
	firstLen := n % k
	if firstLen == 0 {
		firstLen = k
		fullGroups--
	}

	totalLen := n + fullGroups
	res := make([]byte, totalLen)
	writePosition, groupSize, chanCount := totalLen-1, 0, 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '-' {
			continue
		}

		ch := s[i]
		if ch >= 'a' && ch <= 'z' {
			ch -= 'a' - 'A'
		}
		res[writePosition] = ch
		writePosition--
		groupSize++
		chanCount++

		if groupSize == k && chanCount != n {
			res[writePosition] = '-'
			writePosition--
			groupSize = 0
		}
	}
	return string(res)
}

func main() {
	s := "5F3Z-2e-9-w"
	k := 4
	fmt.Println(licenseKeyFormatting(s, k))

	s = "2-5g-3-J"
	k = 2
	fmt.Println(licenseKeyFormatting(s, k))
}
