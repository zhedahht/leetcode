/*
LeetCode 532: https://leetcode.com/problems/k-diff-pairs-in-an-array/
*/

package leetcode

func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}

	numToCount := make(map[int]int)
	for _, num := range nums {
		count := numToCount[num]
		numToCount[num] = count + 1
	}

	pairs := 0
	if k == 0 {
		for _, count := range numToCount {
			if count > 1 {
				pairs++
			}
		}
	} else {
		for num := range numToCount {
			if numToCount[num-k] > 0 {
				pairs++
			}
		}
	}

	return pairs
}
