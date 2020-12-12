/*
LeetCode 341: https://leetcode.com/problems/flatten-nested-list-iterator/
*/

package leetcode

type NestedIterator struct {
	index int
	nums  []int
}

// Constructor341 should be Constructor. Rename it to avoid conflicts
func Constructor341(nestedList []*NestedInteger) *NestedIterator {
	nums := make([]int, 0)
	for _, node := range nestedList {
		nums = append(nums, helper341(node)...)
	}

	return &NestedIterator{
		index: 0,
		nums:  nums,
	}
}

func (this *NestedIterator) Next() int {
	val := this.nums[this.index]
	this.index++
	return val
}

func (this *NestedIterator) HasNext() bool {
	return this.index < len(this.nums)
}

func helper341(nested *NestedInteger) []int {
	result := make([]int, 0)
	if nested.IsInteger() {
		result = append(result, nested.GetInteger())
	} else {
		list := nested.GetList()
		for _, child := range list {
			result = append(result, helper341(child)...)
		}
	}

	return result
}
