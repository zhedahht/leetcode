/*
LeetCode 505: https://leetcode.com/problems/the-maze-ii/description/
*/

package leetcode

import (
	"fmt"
	"math"
)

// The name should be shortestDistance. It's renamed to avoid conflicts.
func shortestDistance505(maze [][]int, start []int, destination []int) int {
	rows, cols := len(maze), len(maze[0])
	stops := make(map[string]int)
	queue := make([][]int, 0)
	queue = append(queue, start)
	stops[encodePos(start)] = 0

	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	found := false
	shortest := math.MaxInt
	for len(queue) > 0 {
		pos := queue[0]
		dist := stops[encodePos(pos)]
		queue = queue[1:]
		if pos[0] == destination[0] && pos[1] == destination[1] {
			found = true
			shortest = min(shortest, dist)
		}

		for _, dir := range dirs {
			i := 1
			for ; true; i++ {
				next := []int{pos[0] + i*dir[0], pos[1] + i*dir[1]}
				if next[0] < 0 || next[0] >= rows || next[1] < 0 || next[1] >= cols ||
					maze[next[0]][next[1]] == 1 {
					break
				}
			}

			if i > 1 {
				next := []int{pos[0] + (i-1)*dir[0], pos[1] + (i-1)*dir[1]}
				nextDist, exists := stops[encodePos(next)]
				if !exists || nextDist > dist+i-1 {
					queue = append(queue, next)
					stops[encodePos(next)] = dist + i - 1
				}
			}
		}
	}

	if found {
		return shortest
	}
	return -1
}

func encodePos(pos []int) string {
	return fmt.Sprintf("%v", pos)
}
