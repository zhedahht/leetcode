/*
LeetCode 399: https://leetcode.com/problems/evaluate-division/
*/

package leetcode

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	getStrToInt := func(equations [][]string) map[string]int {
		strToInt := make(map[string]int)
		for _, e := range equations {
			if _, exists := strToInt[e[0]]; !exists {
				strToInt[e[0]] = len(strToInt)
			}
			if _, exists := strToInt[e[1]]; !exists {
				strToInt[e[1]] = len(strToInt)
			}
		}
		return strToInt
	}

	buildGraph := func(equations [][]string, values []float64, strToInt map[string]int, graph [][]float64) {
		for i := range graph {
			graph[i][i] = 1
		}

		for i, eq := range equations {
			e0, e1 := eq[0], eq[1]
			i1, i2 := strToInt[e0], strToInt[e1]
			graph[i1][i2] = values[i]
			graph[i2][i1] = 1.0 / values[i]
		}
	}

	strToInt := getStrToInt(equations)
	graph := make([][]float64, len(strToInt))
	for i := range graph {
		graph[i] = make([]float64, len(strToInt))
	}
	buildGraph(equations, values, strToInt, graph)

	results := make([]float64, len(queries))
	for i, q := range queries {
		e0, e1 := q[0], q[1]
		if _, exists := strToInt[e0]; !exists {
			results[i] = -1
			continue
		}
		if _, exists := strToInt[e1]; !exists {
			results[i] = -1
			continue
		}

		i0, i1 := strToInt[e0], strToInt[e1]
		visited := make([]bool, len(graph))
		visited[i0] = true
		found, result := helper0399(graph, visited, i0, i1, 1.0)
		if !found {
			result = -1
		}
		results[i] = result
	}
	return results
}

func helper0399(graph [][]float64, visited []bool, src, dst int, result float64) (bool, float64) {
	if src == dst {
		return true, result
	}

	for next, val := range graph[src] {
		if !visited[next] && val != 0 {
			visited[next] = true
			found, newResult := helper0399(graph, visited, next, dst, result*graph[src][next])
			if found {
				return true, newResult
			}
			visited[next] = false
		}
	}
	return false, -1
}
