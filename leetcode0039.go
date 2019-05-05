/*
LeetCode 39: https://leetcode.com/problems/combination-sum/
*/

package leetcode

func combinationSum(candidates []int, target int) [][]int {
	return helper39(candidates, target, 0, make([]int, 0), make([][]int, 0))
}

func helper39(candidates []int, target int, i int, comb []int, result [][]int) [][]int {
	if target == 0 {
		clone := make([]int, len(comb))
		copy(clone, comb)
		result = append(result, clone)
	} else if target > 0 && i < len(candidates) {
		comb = append(comb, candidates[i])
		result = helper39(candidates, target-candidates[i], i, comb, result)
		comb = comb[:len(comb)-1]

		result = helper39(candidates, target, i+1, comb, result)
	}

	return result
}
