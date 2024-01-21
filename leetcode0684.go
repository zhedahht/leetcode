/*
LeetCode 684: https://leetcode.com/problems/redundant-connection/description/
*/

package leetcode

type unionFind struct {
	parent map[int]int
}

func createUnionFind(total int) *unionFind {
	parent := make(map[int]int)
	for i := 1; i <= total; i++ {
		parent[i] = i
	}

	return &unionFind{
		parent: parent,
	}
}

func (u *unionFind) union(num1, num2 int) {
	parent1 := u.findParent(num1)
	parent2 := u.findParent(num2)
	u.parent[parent1] = parent2
}

func (u *unionFind) findParent(num int) int {
	parent := u.parent[num]
	if parent != u.parent[parent] {
		parent = u.findParent(parent)
		u.parent[num] = parent
	}

	return parent
}

func findRedundantConnection(edges [][]int) []int {
	u := createUnionFind(len(edges))
	for _, edge := range edges {
		parent1 := u.findParent(edge[0])
		parent2 := u.findParent(edge[1])
		if parent1 == parent2 {
			return edge
		}

		u.union(edge[0], edge[1])
	}

	// It should never reach here if the input is valid
	return []int{}
}
