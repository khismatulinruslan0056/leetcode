package main

import (
	"fmt"
	"time"
)

/*
You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Example 1:

Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
Example 2:

Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
*/
func climbStairsRec(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStairsRec(n-1) + climbStairsRec(n-2)
}

func climbStairsMap(n int) int {
	mapStairs := make(map[int]int, n)
	mapStairs[1] = 1
	mapStairs[2] = 2
	for i := 3; i <= n; i++ {
		mapStairs[i] = mapStairs[i-1] + mapStairs[i-2]
	}
	return mapStairs[n]
}

func climbStairsSlice(n int) int {
	slStairs := make([]int, n)
	slStairs[0] = 1
	slStairs[1] = 2
	for i := 2; i <= n-1; i++ {
		slStairs[i] = slStairs[i-1] + slStairs[i-2]
	}
	return slStairs[n-1]
}

func climbStairsNum(n int) int {
	if n <= 1 {
		return n
	}
	first, second := 1, 2
	for i := 3; i <= n; i++ {
		first, second = second, first+second
	}
	return second
}

// 1+1+1+1
// 1+1+2
// 1+2+1
// 2+1+1
// 2+2
func main() {
	n := 4
	start := time.Now()
	fmt.Println(climbStairsSlice(n))
	fmt.Println("slice:", time.Since(start))
	start = time.Now()

	fmt.Println(climbStairsMap(n))
	fmt.Println("map:", time.Since(start))
	start = time.Now()
	fmt.Println(climbStairsRec(n))
	fmt.Println("rec:", time.Since(start))

	start = time.Now()
	fmt.Println(climbStairsNum(n))
	fmt.Println("num:", time.Since(start))

	n = 45
	start = time.Now()
	fmt.Println(climbStairsSlice(n))
	fmt.Println("slice:", time.Since(start))
	start = time.Now()

	fmt.Println(climbStairsMap(n))
	fmt.Println("map:", time.Since(start))
	start = time.Now()
	fmt.Println(climbStairsRec(n))
	fmt.Println("rec:", time.Since(start))
	start = time.Now()
	fmt.Println(climbStairsNum(n))
	fmt.Println("num:", time.Since(start))

	//fmt.Println(climbStairs(n))
}
