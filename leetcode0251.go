/*
LeetCode 251: https://leetcode.com/problems/flatten-2d-vector/
*/

package leetcode

type Vector2D struct {
	nums  [][]int
	group int
	index int
}

// Constructor251 should be Constructor. Rename it to avoid conflicts
func Constructor251(v [][]int) Vector2D {
	return Vector2D{
		nums:  v,
		group: 0,
		index: 0,
	}
}

func (this *Vector2D) Next() int {
	this.HasNext()

	val := this.nums[this.group][this.index]
	this.index++
	return val
}

func (this *Vector2D) HasNext() bool {
	for this.group < len(this.nums)-1 && this.index == len(this.nums[this.group]) {
		this.group++
		this.index = 0
	}

	return this.group < len(this.nums) && this.index < len(this.nums[this.group])
}
