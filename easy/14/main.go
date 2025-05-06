package main

import "fmt"

/*Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".



Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.


Constraints:

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of only lowercase English letters if it is non-empty.*/

func longestCommonPrefix(strs []string) string {
	//if len(strs) == 0 {
	//	return ""
	//}
	//if len(strs) == 1 {
	//	return strs[0]
	//}
	//short := shortestString(strs)
	//for i := 0; i < len(strs); i++ {
	//	for len(short) > 0 && !isPrefix(short, strs[i]) {
	//		short = short[:len(short)-1]
	//	}
	//	if short == "" {
	//		return ""
	//	}
	//}
	//return short
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	for i := 0; ; i++ {
		for _, str := range strs {
			if i == len(str) || str[i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
}

func isPrefix(short, str string) bool {
	if len(short) > len(str) {
		return false
	}
	fmt.Println(short, str[:len(short)])
	return short == str[:len(short)]
}

func shortestString(strs []string) string {
	short := strs[0]

	for i := 1; i < len(strs); i++ {
		if len(short) > len(strs[i]) {
			short = strs[i]
		}
	}
	return short
}

func main() {
	var strs []string
	strs = []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs), strs)
	strs = []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strs), strs)

}
