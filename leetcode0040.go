/*
LeetCode 40: https://leetcode.com/problems/combination-sum-ii/
*/

package leetcode

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return helper40(candidates, target, 0, make([]int, 0), make([][]int, 0))
}

func helper40(candidates []int, target int, i int, comb []int, result [][]int) [][]int {
	if target == 0 {
		clone := make([]int, len(comb))
		copy(clone, comb)
		result = append(result, clone)
	} else if target > 0 && i < len(candidates) {
		comb = append(comb, candidates[i])
		result = helper40(candidates, target-candidates[i], i+1, comb, result)
		comb = comb[:len(comb)-1]

		j := i
		for j < len(candidates) && candidates[j] == candidates[i] {
			j++
		}

		result = helper40(candidates, target, j, comb, result)
	}

	return result
}
