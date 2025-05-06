package main

/*
Given an integer rowIndex, return the rowIndexth (0-indexed) row of the Pascal's triangle.

In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:




Example 1:

Input: rowIndex = 3
Output: [1,3,3,1]
Example 2:

Input: rowIndex = 0
Output: [1]
Example 3:

Input: rowIndex = 1
Output: [1,1]*/

import "fmt"

func main() {
	rowIndex := 3
	fmt.Println(getRow(rowIndex))

	rowIndex = 5
	fmt.Println(getRow(rowIndex))

	rowIndex = 1
	fmt.Println(getRow(rowIndex))

}

func getRow(rowIndex int) []int {
	//if rowIndex == 0 {
	//	return []int{1}
	//}
	//res := make([][]int, rowIndex+1)
	//for i := 0; i <= rowIndex; i++ {
	//	res[i] = make([]int, i+1)
	//	res[i][0], res[i][i] = 1, 1
	//	for j := 1; j <= i/2; j++ {
	//		val := res[i-1][j-1] + res[i-1][j]
	//		res[i][j], res[i][i-j] = val, val
	//	}
	//
	//}
	//return res[rowIndex]

	result := make([]int, rowIndex+1)
	result[0] = 1

	if rowIndex == 0 {
		return result
	}

	result[1] = 1
	result[len(result)-1] = 1

	if rowIndex == 1 {
		return result
	}

	for row := 2; row < (rowIndex + 1); row++ {
		for col := row / 2; col >= 1; col-- {
			value := result[col-1] + result[col]
			result[col] = value
			result[row-col] = value
		}
	}

	return result
}
