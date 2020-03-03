/*
LeetCode 532: https://leetcode.com/problems/k-diff-pairs-in-an-array/
*/

package leetcode

func findPairs(nums []int, k int) int {
	numToCount := make(map[int]int)
	for _, num := range nums {
		count := numToCount[num]
		numToCount[num] = count + 1
	}

	pairs := 0
	for num, count := range numToCount {
		if (k == 0 && count > 1) || (k > 0 && numToCount[num-k] > 0) {
			pairs++
		}
	}

	return pairs
}
