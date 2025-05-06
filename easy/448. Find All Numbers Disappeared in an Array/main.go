package main

import (
	"fmt"
)

/*
 Given an array nums of n integers where nums[i] is in the range [1, n], return an array of all the integers in the range [1, n] that do not appear in nums.



Example 1:

Input: nums = [4,3,2,7,8,2,3,1]
Output: [5,6]
Example 2:

Input: nums = [1,1]
Output: [2]


Constraints:

n == nums.length
1 <= n <= 105
1 <= nums[i] <= n


Follow up: Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.*/

func findDisappearedNumbers(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	result := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i]-1 >= 0 {
			if nums[nums[i]-1] > 0 {
				nums[nums[i]-1] *= -1
			}
		} else {
			if nums[-nums[i]-1] > 0 {
				nums[-nums[i]-1] *= -1
			}
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			result = append(result, i+1)
		}
	}

	return result
}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(nums))
	nums = []int{1, 1}
	fmt.Println(findDisappearedNumbers(nums))
}
