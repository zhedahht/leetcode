/*
LeetCode 170: https://leetcode.com/problems/two-sum-iii-data-structure-design/
*/

package leetcode

import "sort"

type TwoSum struct {
	nums []int
}

//NOTE: The function name should be Constructor. Rename it to avoid conflicts.
/** Initialize your data structure here. */
func Constructor170() TwoSum {
	return TwoSum{}
}

/** Add the number to an internal data structure.. */
func (this *TwoSum) Add(number int) {
	i := sort.SearchInts(this.nums, number)
	this.nums = append(this.nums, number)
	copy(this.nums[i+1:], this.nums[i:len(this.nums)-1])
	this.nums[i] = number
}

/** Find if there exists any pair of numbers which sum is equal to the value. */
func (this *TwoSum) Find(value int) bool {
	for i, j := 0, len(this.nums)-1; i < j; {
		if this.nums[i]+this.nums[j] == value {
			return true
		} else if this.nums[i]+this.nums[j] > value {
			j--
		} else {
			i++
		}
	}

	return false
}

/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */
