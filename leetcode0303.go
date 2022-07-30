/*
LeetCode 303: https://leetcode.com/problems/range-sum-query-immutable/
*/

package leetcode

type NumArray struct {
	sums []int
}

// Constructor0303 should be Constructor. It's renamed to avoid conflicts.
func Constructor0303(nums []int) NumArray {
	sums := make([]int, len(nums)+1)
	sums[0] = 0
	for i, num := range nums {
		sums[i+1] = sums[i] + num
	}
	return NumArray{sums: sums}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sums[right+1] - this.sums[left]
}
