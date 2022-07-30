/*
LeetCode 384: https://leetcode.com/problems/shuffle-an-array/
*/

package leetcode

import "math/rand"

type Solution struct {
	nums []int
}

// Constructor0384 should be Constructor. Rename it to avoid name conflicts.
func Constructor0384(nums []int) Solution {
	return Solution{
		nums: nums,
	}
}

func (this *Solution) Reset() []int {
	return this.nums
}

func (this *Solution) Shuffle() []int {
	clone := make([]int, len(this.nums))
	copy(clone, this.nums)

	for i := len(clone) - 1; i >= 0; i-- {
		r := rand.Intn(i + 1)
		clone[i], clone[r] = clone[r], clone[i]
	}
	return clone
}
