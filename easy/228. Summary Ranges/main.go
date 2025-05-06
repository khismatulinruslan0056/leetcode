package main

import (
	"fmt"
	"strconv"
)

/*
 You are given a sorted unique integer array nums.

A range [a,b] is the set of all integers from a to b (inclusive).

Return the smallest sorted list of ranges that cover all the numbers in the array exactly. That is, each element of nums is covered by exactly one of the ranges, and there is no integer x such that x is in one of the ranges but not in nums.

Each range [a,b] in the list should be output as:

"a->b" if a != b
"a" if a == b


Example 1:

Input: nums = [0,1,2,4,5,7]
Output: ["0->2","4->5","7"]
Explanation: The ranges are:
[0,2] --> "0->2"
[4,5] --> "4->5"
[7,7] --> "7"
Example 2:

Input: nums = [0,2,3,4,6,8,9]
Output: ["0","2->4","6","8->9"]
Explanation: The ranges are:
[0,0] --> "0"
[2,4] --> "2->4"
[6,6] --> "6"
[8,9] --> "8->9"


Constraints:

0 <= nums.length <= 20
-231 <= nums[i] <= 231 - 1
All the values of nums are unique.
nums is sorted in ascending order.*/

func summaryRanges(nums []int) []string {
	//if len(nums) < 1 {
	//	return nil
	//}
	//if len(nums) == 1 {
	//	return []string{strconv.Itoa(nums[0])}
	//}
	//result := make([]string, 0, len(nums))
	//start := nums[0]
	//indStart := 0
	//for i := 1; i < len(nums); i++ {
	//	str := ""
	//	if nums[i]-nums[i-1] > 1 {
	//		if i == indStart+1 {
	//			str = "\"" + strconv.Itoa(start) + "\""
	//		} else {
	//			str = "\"" + strconv.Itoa(start) + "->" + strconv.Itoa(nums[i-1]) + "\""
	//		}
	//		result = append(result, str)
	//		start = nums[i]
	//		indStart = i
	//	}
	//	if i == len(nums)-1 {
	//		if i == indStart {
	//			str = "\"" + strconv.Itoa(start) + "\""
	//		} else {
	//			str = "\"" + strconv.Itoa(start) + "->" + strconv.Itoa(nums[i]) + "\""
	//		}
	//		result = append(result, str)
	//	}
	//}
	//return result
	if len(nums) == 0 {
		return nil
	}

	result := make([]string, 0)
	start := nums[0]

	for i := 1; i <= len(nums); i++ {
		// Если текущий элемент не продолжает последовательность или мы достигли конца
		if i == len(nums) || nums[i]-nums[i-1] > 1 {
			if start == nums[i-1] {
				result = append(result, strconv.Itoa(start))
			} else {
				result = append(result, strconv.Itoa(start)+"->"+strconv.Itoa(nums[i-1]))
			}
			if i < len(nums) {
				start = nums[i]
			}
		}
	}

	return result
}

func main() {
	nums := []int{0, 1, 2, 4, 5, 7}
	fmt.Println("[\"0->2\",\"4->5\",\"7\"]")
	fmt.Println(summaryRanges(nums))
	nums = []int{0, 2, 3, 4, 6, 8, 9}
	fmt.Println("[\"0\",\"2->4\",\"6\",\"8->9\"]")
	fmt.Println(summaryRanges(nums))
	nums = []int{-1}
	fmt.Println("[\"0\",\"2->4\",\"6\",\"8->9\"]")
	fmt.Println(summaryRanges(nums))
}
