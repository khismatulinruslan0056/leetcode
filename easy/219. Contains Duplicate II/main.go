package main

import "fmt"

/*
 Given an integer array nums and an integer k, return true if there are two distinct indices i and j
in the array such that nums[i] == nums[j] and abs(i - j) <= k.



Example 1:

Input: nums = [1,2,3,1], k = 3
Output: true
Example 2:

Input: nums = [1,0,1,1], k = 1
Output: true
Example 3:

Input: nums = [1,2,3,1,2,3], k = 2
Output: false


Constraints:

1 <= nums.length <= 105
-109 <= nums[i] <= 109
0 <= k <= 105
Seen this question in a real interview before?
1/5
*/

func containsNearbyDuplicate(nums []int, k int) bool {
	mapNearbyDuplicate := make(map[int]int)
	for i, num := range nums {
		if _, ok := mapNearbyDuplicate[num]; ok {
			if i-mapNearbyDuplicate[num] <= k {
				return true
			}
		}
		mapNearbyDuplicate[num] = i

	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	k := 3
	fmt.Println(containsNearbyDuplicate(nums, k))

	nums = []int{1, 0, 1, 1}
	k = 1
	fmt.Println(containsNearbyDuplicate(nums, k))

	nums = []int{1, 2, 3, 1, 2, 3}
	k = 2
	fmt.Println(containsNearbyDuplicate(nums, k))

}
