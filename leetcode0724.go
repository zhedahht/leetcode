/*
LeetCode 724: https://leetcode.com/problems/find-pivot-index/
*/

package leetcode

func pivotIndex(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	sum := 0
	for i, num := range nums {
		sum += num
		if sum == total-sum+num {
			return i
		}
	}

	return -1
}
