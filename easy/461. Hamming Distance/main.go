package main

import "fmt"

/*
 The Hamming distance between two integers is the number of positions at which the corresponding bits are different.

Given two integers x and y, return the Hamming distance between them.



Example 1:

Input: x = 1, y = 4
Output: 2
Explanation:
1   (0 0 0 1)
4   (0 1 0 0)
       ↑   ↑
The above arrows point to positions where the corresponding bits are different.
Example 2:

Input: x = 3, y = 1
Output: 1
1   (0 0 0 1)
3   (0 0 1 1)
         ↑

Constraints:

0 <= x, y <= 231 - 1


Note: This question is the same as 2220: Minimum Bit Flips to Convert Number.*/

func hammingDistance(x int, y int) int {
	count := 0

	for x > 0 || y > 0 {
		if x%2 != y%2 {
			count++
		}
		x /= 2
		y /= 2
	}

	return count
}

func main() {
	x := 1
	y := 4
	fmt.Println(hammingDistance(x, y))

	x = 3
	y = 1
	fmt.Println(hammingDistance(x, y))
}
