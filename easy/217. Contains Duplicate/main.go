package main

import "fmt"

/*
 Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.



Example 1:

Input: nums = [1,2,3,1]

Output: true

Explanation:

The element 1 occurs at the indices 0 and 3.

Example 2:

Input: nums = [1,2,3,4]

Output: false

Explanation:

All elements are distinct.

Example 3:

Input: nums = [1,1,1,3,3,4,3,2,4,2]

Output: true



Constraints:

1 <= nums.length <= 105
-109 <= nums[i] <= 109
Seen this question in a real interview before?*/

func containsDuplicate(nums []int) bool {
	//mapCounter := make(map[int]struct{})
	//for _, num := range nums {
	//	if _, ok := mapCounter[num]; ok {
	//		return true
	//	}
	//	mapCounter[num] = struct{}{}
	//}
	//return false

	if len(nums) < 2 {
		return false
	}
	isSorted := false

	for !isSorted {
		isSorted = true
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] == nums[i+1] {
				return true
			}
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				isSorted = false
			}
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(nums))
	nums = []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(nums))
	nums = []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(containsDuplicate(nums))
}
