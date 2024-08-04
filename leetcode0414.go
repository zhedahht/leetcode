/*
LeetCode 414: https://leetcode.com/problems/third-maximum-number/description/
*/

package leetcode

import "container/heap"

func thirdMax(nums []int) int {
	numsCount := make(map[int]int)
	for _, num := range nums {
		numsCount[num]++
	}

	h := &MinHeap{}
	heap.Init(h)

	for num := range numsCount {
		if h.Len() < 3 {
			heap.Push(h, num)
		} else if (*h)[0] < num {
			heap.Pop(h)
			heap.Push(h, num)
		}
	}

	if h.Len() == 2 {
		heap.Pop(h)
	}

	return (*h)[0]
}
