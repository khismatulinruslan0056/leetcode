package main

import "fmt"

/*
Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.

Example 1:

Input: nums = [1,3,5,6], target = 5
Output: 2
Example 2:

Input: nums = [1,3,5,6], target = 2
Output: 1
Example 3:

Input: nums = [1,3,5,6], target = 7
Output: 4

Constraints:

1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums contains distinct values sorted in ascending order.
-104 <= target <= 104
*/
func searchInsert(nums []int, target int) int {
	//i := 0
	//for i = 0; i < len(nums); i++ {
	//	if nums[i] >= target {
	//		return i
	//	}
	//}
	//return i
	//first := 0
	//last := len(nums) - 1
	//if nums[first] >= target {
	//	return 0
	//}
	//if nums[last] < target {
	//	return len(nums)
	//}
	//mid := (first + last) / 2
	//for first <= last {
	//	if nums[mid] == target {
	//		return mid
	//	}
	//	if nums[mid] < target {
	//		first = mid
	//	} else {
	//		last = mid
	//	}
	//	if last-first < 2 {
	//		return mid - 1
	//	}
	//	mid = (first + last) / 2
	//}
	//return mid
	first := 0
	last := len(nums) - 1
	if nums[first] >= target {
		return 0
	}

	for first <= last {
		mid := (first + last) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			first = mid + 1
		} else {
			last = mid - 1
		}
	}
	return first
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 5
	fmt.Println(searchInsert(nums, target))
	nums = []int{1, 3, 5, 6}
	target = 2
	fmt.Println(searchInsert(nums, target))
	nums = []int{1, 3, 5, 6}
	target = 7
	fmt.Println(searchInsert(nums, target))
}
