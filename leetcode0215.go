/*
LeetCode 215: https://leetcode.com/problems/kth-largest-element-in-an-array/
*/

package leetcode

import "container/heap"

func findKthLargest(nums []int, k int) int {
	minHeap := &MinHeap{}
	for _, num := range nums {
		if minHeap.Len() < k {
			heap.Push(minHeap, num)
		} else {
			top := (*minHeap)[0]
			if top < num {
				heap.Pop(minHeap)
				heap.Push(minHeap, num)
			}
		}
	}

	return (*minHeap)[0]
}
