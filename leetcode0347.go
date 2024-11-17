/*
LeetCode 347: https://leetcode.com/problems/top-k-frequent-elements/description/
*/

package leetcode

import "container/heap"

type MinHeap347 [][]int

func (h MinHeap347) Len() int           { return len(h) }
func (h MinHeap347) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h MinHeap347) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap347) Push(x any) {
	*h = append(*h, x.([]int))
}

func (h *MinHeap347) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// The function name should be topKFrequent. It's renamed to avoid conflicts.
func topKFrequent347(nums []int, k int) []int {
	numCounts := make(map[int]int)
	for _, num := range nums {
		numCounts[num]++
	}

	h := &MinHeap347{}
	heap.Init(h)

	for num, count := range numCounts {
		if h.Len() < k {
			heap.Push(h, []int{num, count})
		} else {
			top := (*h)[0]
			if top[1] < count {
				heap.Pop(h)
				heap.Push(h, []int{num, count})
			}
		}
	}

	result := make([]int, 0)
	for i := range [][]int(*h) {
		result = append(result, [][]int(*h)[i][0])
	}

	return result
}
