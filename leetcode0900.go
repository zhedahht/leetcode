/*
LeetCode 900: https://leetcode.com/problems/rle-iterator/
*/

package leetcode

type RLEIterator struct {
	nums []int
}

// Constructor900 should be Constructor.
// Rename it to avoid name conflicts.
func Constructor900(A []int) RLEIterator {
	return RLEIterator{nums: A}
}

func (this *RLEIterator) Next(n int) int {
	for len(this.nums) > 0 {
		if this.nums[0] >= n {
			this.nums[0] -= n
			val := this.nums[1]
			if this.nums[0] == 0 {
				this.nums = this.nums[2:]
			}

			return val
		}

		n -= this.nums[0]
		this.nums = this.nums[2:]
	}

	return -1
}
