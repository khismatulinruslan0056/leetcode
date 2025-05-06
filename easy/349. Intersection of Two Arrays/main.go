package main

import "fmt"

/*
 Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must be unique and you may return the result in any order.



Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]
Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]
Explanation: [4,9] is also accepted.


Constraints:

1 <= nums1.length, nums2.length <= 1000
0 <= nums1[i], nums2[i] <= 1000*/

func intersection(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)
	seen := make(map[int]bool)
	for i := 0; i < len(nums1); i++ {
		seen[nums1[i]] = true
	}
	for i := 0; i < len(nums2); i++ {
		if seen[nums2[i]] {
			result = append(result, nums2[i])
			seen[nums2[i]] = false
		}
	}
	return result
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersection(nums1, nums2))
	nums1 = []int{4, 9, 5}
	nums2 = []int{9, 4, 9, 8, 4}
	fmt.Println(intersection(nums1, nums2))
}
