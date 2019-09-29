/*
LeetCode 167: https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
*/

package leetcode

func twoSum(numbers []int, target int) []int {
	for i, j := 0, len(numbers)-1; i < j; {
		if numbers[i]+numbers[j] == target {
			return []int{i + 1, j + 1}
		}

		if numbers[i]+numbers[j] > target {
			j--
		} else {
			i++
		}
	}

	return []int{-1, -1}
}
