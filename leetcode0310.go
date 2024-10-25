/*
LeetCode 310: https://leetcode.com/problems/minimum-height-trees/description/
*/

package leetcode

func findMinHeightTrees(n int, edges [][]int) []int {
	graph := make(map[int][]int)
	inOrders := make(map[int]int)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, 0)
		inOrders[i] = 0
	}

	for _, edge := range edges {
		in, out := edge[0], edge[1]
		inOrders[in]++
		inOrders[out]++

		graph[in] = append(graph[in], out)
		graph[out] = append(graph[out], in)
	}

	queue1, queue2 := make([]int, 0), make([]int, 0)
	visited := make(map[int]bool)
	for n, inOrder := range inOrders {
		if inOrder <= 1 {
			queue1 = append(queue1, n)
			visited[n] = true
		}
	}

	last := make([]int, 0)
	for len(queue1) > 0 {
		n := queue1[0]
		queue1 = queue1[1:]

		last = append(last, n)
		for _, next := range graph[n] {
			if !visited[next] {
				inOrders[next]--
				if inOrders[next] <= 1 {
					queue2 = append(queue2, next)
					visited[next] = true
				}
			}
		}

		if len(queue1) == 0 {
			queue1 = queue2
			queue2 = make([]int, 0)
			if len(queue1) > 0 {
				last = make([]int, 0)
			}
		}
	}

	return last
}
