/*
LeetCode 332: https://leetcode.com/problems/reconstruct-itinerary/
*/

package leetcode

import "sort"

type cityWithCount struct {
	city  string
	count int
}

func findItinerary(tickets [][]string) []string {
	cityGraph := make(map[string]map[string]int)
	for _, t := range tickets {
		if _, exists := cityGraph[t[0]]; !exists {
			cityGraph[t[0]] = make(map[string]int)
		}
		cityGraph[t[0]][t[1]]++
	}

	graph := make(map[string][]*cityWithCount)
	for city, nextCount := range cityGraph {
		nexts := make([]*cityWithCount, 0)
		for c, count := range nextCount {
			nexts = append(nexts, &cityWithCount{
				city:  c,
				count: count,
			})
		}

		sort.Slice(nexts, func(i, j int) bool {
			return nexts[i].city < nexts[j].city
		})
		graph[city] = nexts
	}

	return helper0332(graph, "JFK", []string{}, len(tickets))
}

func helper0332(graph map[string][]*cityWithCount, start string, path []string, length int) []string {
	path = append(path, start)
	if len(path) == length+1 {
		return path
	}

	nexts := graph[start]
	for _, next := range nexts {
		if next.count == 0 {
			continue
		}

		next.count--
		path = helper0332(graph, next.city, path, length)
		if len(path) == length+1 {
			return path
		}
		next.count++
	}

	path = path[:len(path)-1]
	return path
}
