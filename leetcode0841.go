/*
LeetCode 841: https://leetcode.com/problems/keys-and-rooms/description/
*/

package leetcode

func canVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms))
	visited[0] = true
	queue := []int{0}
	for len(queue) > 0 {
		room := queue[0]
		queue = queue[1:]

		for _, next := range rooms[room] {
			if !visited[next] {
				visited[next] = true
				queue = append(queue, next)
			}
		}
	}

	for _, flag := range visited {
		if !flag {
			return false
		}
	}

	return true
}
