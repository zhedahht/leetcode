/*
LeetCode 286: https://leetcode.com/problems/walls-and-gates/description/
*/

package leetcode

func wallsAndGates(rooms [][]int) {
	rows, cols := len(rooms), len(rooms[0])
	queue1, queue2 := make([][]int, 0), make([][]int, 0)
	for i, row := range rooms {
		for j, val := range row {
			if val == 0 {
				queue1 = append(queue1, []int{i, j})
			}
		}
	}

	dist := 0
	dirs := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	for len(queue1) > 0 {
		pos := queue1[0]
		queue1 = queue1[1:]

		for _, dir := range dirs {
			next := []int{pos[0] + dir[0], pos[1] + dir[1]}
			if next[0] < 0 || next[0] >= rows || next[1] < 0 || next[1] >= cols {
				continue
			}

			if rooms[next[0]][next[1]] < 0 || rooms[next[0]][next[1]] <= dist+1 {
				continue
			}

			queue2 = append(queue2, next)
			rooms[next[0]][next[1]] = dist + 1
		}

		if len(queue1) == 0 {
			queue1 = queue2
			queue2 = make([][]int, 0)
			dist++
		}
	}
}
