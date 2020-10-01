/*
LeetCode 133: https://leetcode.com/problems/clone-graph/
*/

package leetcode

// Node133 should be Node. Rename it to avoid conflict
type Node133 struct {
	Val       int
	Neighbors []*Node133
}

func cloneGraph(node *Node133) *Node133 {
	if node == nil {
		return nil
	}

	m := make(map[*Node133]*Node133)
	queue := make([]*Node133, 0)
	visited := make(map[*Node133]bool)

	c := &Node133{
		Val: node.Val,
	}

	queue = append(queue, node)
	visited[node] = true
	m[node] = c

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		cloneN := m[n]

		for _, neighbor := range n.Neighbors {
			if !visited[neighbor] {
				cloneNeighbor := &Node133{
					Val: neighbor.Val,
				}

				queue = append(queue, neighbor)
				m[neighbor] = cloneNeighbor
				visited[neighbor] = true
			}

			cloneN.Neighbors = append(cloneN.Neighbors, m[neighbor])
		}
	}

	return c
}
