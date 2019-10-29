/*
LeetCode 870: https://leetcode.com/problems/advantage-shuffle/
*/

package leetcode

import "sort"

func advantageCount(A []int, B []int) []int {
	result := make([]int, len(A))
	sort.Ints(A)

	bWithIndex := make([][]int, len(B))
	for i, b := range B {
		bWithIndex[i] = []int{b, i}
	}

	sort.Slice(bWithIndex, func(i, j int) bool {
		return bWithIndex[i][0] < bWithIndex[j][0]
	})

	i, j := 0, len(A)-1
	for k := len(B) - 1; k >= 0; k-- {
		a, b := A[j], bWithIndex[k][0]
		if a > b {
			result[bWithIndex[k][1]] = a
			j--
		} else {
			result[bWithIndex[k][1]] = A[i]
			i++
		}
	}

	return result
}
