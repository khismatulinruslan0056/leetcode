package main

import "fmt"

/*
 Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.



Example 1:

Input: nums = [3,2,3]
Output: 3
Example 2:

Input: nums = [2,2,1,1,1,2,2]
Output: 2


Constraints:

n == nums.length
1 <= n <= 5 * 104
-109 <= nums[i] <= 109


Follow-up: Could you solve the problem in linear time and in O(1) space?*/

func main() {
	nums := []int{3, 2, 3}
	fmt.Println(3, nums, majorityElement(nums))

	nums = []int{2, 2, 1, 1, 1, 2, 2, 3, 4, 4, 7, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5}
	fmt.Println(2, nums, majorityElement(nums))
}

func majorityElement(nums []int) int {
	//mapCount := make(map[int]int)
	//for _, num := range nums {
	//	mapCount[num]++
	//}
	//major := nums[0]
	//majorityElement := mapCount[major]
	//for num, count := range mapCount {
	//	if major < count {
	//		major = count
	//		majorityElement = num
	//	}
	//}
	//return majorityElement
	candidate := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			candidate = nums[i]
			count = 1
		} else if nums[i] == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}
