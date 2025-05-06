package main

import "fmt"

/*
 You are given row x col grid representing a map where grid[i][j] = 1 represents land and grid[i][j] = 0 represents water.

Grid cells are connected horizontally/vertically (not diagonally). The grid is completely surrounded by water, and there is exactly one island (i.e., one or more connected land cells).

The island doesn't have "lakes", meaning the water inside isn't connected to the water around the island. One cell is a square with side length 1. The grid is rectangular, width and height don't exceed 100. Determine the perimeter of the island.



Example 1:


Input: grid = [[0,1,0,0],[1,1,1,0],[0,1,0,0],[1,1,0,0]]
Output: 16
Explanation: The perimeter is the 16 yellow stripes in the image above.
Example 2:

Input: grid = [[1]]
Output: 4
Example 3:

Input: grid = [[1,0]]
Output: 4


Constraints:

row == grid.length
col == grid[i].length
1 <= row, col <= 100
grid[i][j] is 0 or 1.
There is exactly one island in grid.*/

func islandPerimeter(grid [][]int) int {
	perimeter := 0
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	columns := len(grid[0])
	borders := make([]bool, columns)
	for i := 0; i < len(grid); i++ {
		flag := false
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				perimeter += 4
				if borders[j] {
					perimeter -= 2
				}
				if flag {
					perimeter -= 2
				}
				borders[j] = true
				flag = true
			} else {
				borders[j] = false
				flag = false
			}
		}
	}
	return perimeter
}

func main() {
	grid := [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}
	fmt.Println(islandPerimeter(grid))

	grid = [][]int{{1}}
	fmt.Println(islandPerimeter(grid))

	grid = [][]int{{1, 0}}
	fmt.Println(islandPerimeter(grid))
}
