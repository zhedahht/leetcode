/*
LeetCode 994: https://leetcode.com/problems/rotting-oranges/
*/

package leetcode

func orangesRotting(grid [][]int) int {
	queue1, queue2 := make([][]int, 0), make([][]int, 0)
	total, rotten := 0, 0
	for r, row := range grid {
		for c, v := range row {
			if v == 2 {
				queue1 = append(queue1, []int{r, c})
				total, rotten = total+1, rotten+1
			} else if v == 1 {
				total++
			}
		}
	}

	if total == 0 {
		return 0
	}

	days := -1
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for len(queue1) > 0 {
		pos := queue1[0]
		queue1 = queue1[1:]

		for _, dir := range dirs {
			r, c := dir[0]+pos[0], dir[1]+pos[1]
			if r >= 0 && r < len(grid) &&
				c >= 0 && c < len(grid[0]) &&
				grid[r][c] == 1 {
				grid[r][c] = 2
				rotten++
				queue2 = append(queue2, []int{r, c})
			}
		}

		if len(queue1) == 0 {
			queue1 = queue2
			queue2 = make([][]int, 0)
			days++
		}
	}

	if rotten != total {
		return -1
	}

	return days
}
