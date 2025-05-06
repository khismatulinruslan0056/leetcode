package main

import (
	"fmt"
	"math"
)

/*
 Given an integer array nums, return the third distinct maximum number in this array. If the third maximum does not exist, return the maximum number.



Example 1:

Input: nums = [3,2,1]
Output: 1
Explanation:
The first distinct maximum is 3.
The second distinct maximum is 2.
The third distinct maximum is 1.
Example 2:

Input: nums = [1,2]
Output: 2
Explanation:
The first distinct maximum is 2.
The second distinct maximum is 1.
The third distinct maximum does not exist, so the maximum (2) is returned instead.
Example 3:

Input: nums = [2,2,3,1]
Output: 1
Explanation:
The first distinct maximum is 3.
The second distinct maximum is 2 (both 2's are counted together since they have the same value).
The third distinct maximum is 1.


Constraints:

1 <= nums.length <= 104
-231 <= nums[i] <= 231 - 1


Follow up: Can you find an O(n) solution?*/

func thirdMax(nums []int) int {
	firstmax, secondmax, thirdmax := math.MinInt, math.MinInt, math.MinInt
	for _, num := range nums {
		if firstmax < num {
			thirdmax = secondmax
			secondmax = firstmax
			firstmax = num
		}
		if secondmax < num && firstmax != num {
			thirdmax = secondmax
			secondmax = num
		}
		if thirdmax < num && secondmax != num && firstmax != num {
			thirdmax = num
		}
	}
	if len(nums) < 3 {
		return firstmax
	}
	return thirdmax
}

func main() {
	nums := []int{3, 2, 1}
	fmt.Println(nums, thirdMax(nums))
	nums = []int{1, 2}
	fmt.Println(nums, thirdMax(nums))
	nums = []int{2, 2, 3, 1}
	fmt.Println(nums, thirdMax(nums))
}
