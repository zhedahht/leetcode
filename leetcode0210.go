/*
LeetCode 210: https://leetcode.com/problems/course-schedule-ii/
*/

package leetcode

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make(map[int][]int)
	inCounts := make(map[int]int)
	for i := 0; i < numCourses; i++ {
		graph[i] = make([]int, 0)
		inCounts[i] = 0
	}

	for _, pair := range prerequisites {
		graph[pair[1]] = append(graph[pair[1]], pair[0])
		inCounts[pair[0]]++
	}

	queue := make([]int, 0)
	for k, count := range inCounts {
		if count == 0 {
			queue = append(queue, k)
		}
	}

	order := make([]int, 0)
	for len(queue) > 0 {
		course := queue[0]
		order = append(order, course)
		queue = queue[1:]

		for _, neighbor := range graph[course] {
			inCounts[neighbor]--

			if inCounts[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	if len(order) < numCourses {
		return make([]int, 0)
	}

	return order
}
