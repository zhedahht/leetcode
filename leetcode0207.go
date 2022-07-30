/*
LeetCode 207: https://leetcode.com/problems/course-schedule/
*/

package leetcode

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int]map[int]bool)
	in := make(map[int]int)
	for i := 0; i < numCourses; i++ {
		graph[i] = make(map[int]bool)
		in[i] = 0
	}
	for _, prer := range prerequisites {
		graph[prer[1]][prer[0]] = true
		in[prer[0]]++
	}

	queue := make([]int, 0)
	for c, val := range in {
		if val == 0 {
			queue = append(queue, c)
		}
	}

	for len(queue) != 0 {
		c := queue[0]
		queue = queue[1:]

		for next := range graph[c] {
			in[next]--
			if in[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	for _, val := range in {
		if val > 0 {
			return false
		}
	}
	return true
}
