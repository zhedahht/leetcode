/*
LeetCode 815: https://leetcode.com/problems/bus-routes/
*/

package leetcode

func numBusesToDestination(routes [][]int, source int, target int) int {
	buildGraph := func(routes [][]int) (map[int]map[int]bool, map[int]map[int]bool) {
		stopToRoutes := make(map[int]map[int]bool)
		routeToStops := make(map[int]map[int]bool)
		for i, route := range routes {
			routeToStops[i] = make(map[int]bool)
			for _, stop := range route {
				routeToStops[i][stop] = true

				if _, exists := stopToRoutes[stop]; !exists {
					stopToRoutes[stop] = make(map[int]bool)
				}
				stopToRoutes[stop][i] = true
			}
		}
		return stopToRoutes, routeToStops
	}
	
	stopToRoutes, routeToStops := buildGraph(routes)
	visitedStop := make(map[int]bool)
	visitedRoute := make(map[int]bool)
	queue1, queue2 := make([]int, 0), make([]int, 0)
	queue1 = append(queue1, source)
	visitedStop[source] = true
	steps := 0
	for len(queue1) > 0 {
		stop := queue1[0]
		queue1 = queue1[1:]
		if stop == target {
			return steps
		}

		rs := stopToRoutes[stop]
		for r := range rs {
			if visitedRoute[r] {
				continue
			}
			visitedRoute[r] = true
			nexts := routeToStops[r]
			for next := range nexts {
				if !visitedStop[next] {
					queue2 = append(queue2, next)
					visitedStop[next] = true
				}
			}
		}

		if len(queue1) == 0 {
			queue1 = queue2
			queue2 = make([]int, 0)
			steps++
		}
	}
	return -1
}
