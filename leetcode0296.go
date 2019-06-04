/*
LeetCode 296: https://leetcode.com/problems/best-meeting-point/
*/

package leetcode

import (
	"math"
	"sort"
)

func minTotalDistance(grid [][]int) int {
	rows, cols := make([]int, 0), make([]int, 0)
	for i := range grid {
		for j, val := range grid[i] {
			if val == 1 {
				rows = append(rows, i)
				cols = append(cols, j)
			}
		}
	}

	return minDist1D(rows) + minDist1D(cols)
}

func minDist1D(pos []int) int {
	sort.Ints(pos)
	mid := len(pos) / 2
	sum := 0
	for _, val := range pos {
		sum += int(math.Abs(float64(pos[mid]) - float64(val)))
	}

	return sum
}
