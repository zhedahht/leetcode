/*
LeetCode 892: https://leetcode.com/problems/surface-area-of-3d-shapes/
*/

package leetcode

func surfaceArea(grid [][]int) int {
	max := func(m, n int) int {
		if m > n {
			return m
		}

		return n
	}

	result := 0
	if len(grid) == 0 || len(grid[0]) == 0 {
		return result
	}

	for i, row := range grid {
		for j, num := range row {
			left, up, right, down := num, num, num, num
			if i != 0 {
				up = max(0, num-grid[i-1][j])
			}

			if i != len(grid)-1 {
				down = max(0, num-grid[i+1][j])
			}

			if j != 0 {
				left = max(0, num-row[j-1])
			}

			if j != len(row)-1 {
				right = max(0, num-row[j+1])
			}

			result += left + up + right + down
			if num != 0 {
				result += 2
			}
		}
	}

	return result
}
