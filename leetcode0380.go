/*
LeetCode 380: https://leetcode.com/problems/insert-delete-getrandom-o1/description/
*/

package leetcode

import "math/rand"

type RandomizedSet struct {
	nums      []int
	positions map[int]int
}

// The function name should be Constructor. It's renamed to avoid conflicts.
func Constructor380() RandomizedSet {
	return RandomizedSet{
		nums:      make([]int, 0),
		positions: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, exists := this.positions[val]; exists {
		return false
	}

	this.positions[val] = len(this.nums)
	this.nums = append(this.nums, val)

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, exists := this.positions[val]; !exists {
		return false
	}

	pos := this.positions[val]
	if pos != len(this.nums)-1 {
		this.nums[pos] = this.nums[len(this.nums)-1]
		this.positions[this.nums[pos]] = pos
	}

	delete(this.positions, val)
	this.nums = this.nums[:len(this.nums)-1]

	return true
}

func (this *RandomizedSet) GetRandom() int {
	index := rand.Intn(len(this.nums))
	return this.nums[index]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
