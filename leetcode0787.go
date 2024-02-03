/*
LeetCode 787: https://leetcode.com/problems/cheapest-flights-within-k-stops/description/
*/

package leetcode

import "math"

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	graph := buildGraph787(flights)
	minCosts := make(map[int]int)
	for i := 0; i < n; i++ {
		minCosts[i] = math.MaxInt
	}

	minCosts[src] = 0
	queue1 := [][2]int{{src, 0}}
	queue2 := [][2]int{}
	stops := 0
	cheapestPrice := math.MaxInt
	for len(queue1) > 0 {
		city := queue1[0]
		queue1 = queue1[1:]
		minCosts[city[0]] = min(minCosts[city[0]], city[1])
		if city[0] == dst {
			cheapestPrice = min(cheapestPrice, minCosts[city[0]])
		}

		for _, next := range graph[city[0]] {
			if minCosts[city[0]]+next[1] < cheapestPrice && minCosts[city[0]]+next[1] < minCosts[next[0]] {
				queue2 = append(queue2, [2]int{next[0], minCosts[city[0]] + next[1]})
			}
		}

		if len(queue1) == 0 {
			queue1 = queue2
			queue2 = [][2]int{}
			stops++
			if stops > k+1 {
				break
			}
		}
	}

	if cheapestPrice == math.MaxInt {
		return -1
	}

	return cheapestPrice
}

func buildGraph787(flights [][]int) map[int][][2]int {
	graph := make(map[int][][2]int)
	for _, flight := range flights {
		if _, exists := graph[flight[0]]; !exists {
			graph[flight[0]] = make([][2]int, 0)
		}

		graph[flight[0]] = append(graph[flight[0]], [2]int{flight[1], flight[2]})
	}

	return graph
}
