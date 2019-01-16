/*
LeetCode 254: https://leetcode.com/problems/factor-combinations/
*/

package leetcode

import "math"

func getFactors(n int) [][]int {
	_, combinations := getFactorCombination(n, 2, make([]int, 0), make([][]int, 0))
	return combinations
}

func getFactorCombination(n int, minFactor int, combination []int, combinations [][]int) ([]int, [][]int) {
	if n > 1 {
		for i := minFactor; i <= int(math.Sqrt(float64(n))); i++ {
			if n%i == 0 {
				combination = append(combination, i)
				combination = append(combination, n/i)
				temp := make([]int, len(combination))
				copy(temp, combination)
				combinations = append(combinations, temp)
				combination = combination[:len(combination)-1]

				combination, combinations = getFactorCombination(n/i, i, combination, combinations)
				combination = combination[:len(combination)-1]
			}
		}
	}

	return combination, combinations
}
