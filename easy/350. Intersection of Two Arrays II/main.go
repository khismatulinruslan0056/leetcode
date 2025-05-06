package main

import "fmt"

/*
Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must appear as many times as it shows in both arrays and you may return the result in any order.

Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2,2]
Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [4,9]
Explanation: [9,4] is also accepted.

Constraints:

1 <= nums1.length, nums2.length <= 1000
0 <= nums1[i], nums2[i] <= 1000
*/
func intersect(nums1 []int, nums2 []int) []int {
	inter1 := make(map[int]int)
	inter2 := make(map[int]int)

	for _, n := range nums1 {
		inter1[n]++
	}
	for _, n := range nums2 {
		inter2[n]++
	}
	result := make([]int, 0)
	for _, n := range nums1 {
		if v, ok := inter2[n]; ok && v > 0 {
			result = append(result, n)
			inter2[n]--
			inter1[n]--
		}
	}

	return result
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersect(nums1, nums2))
	nums1 = []int{4, 9, 5}
	nums2 = []int{9, 4, 9, 8, 4}
	fmt.Println(intersect(nums1, nums2))
}
