/*
LeetCode 315: https://leetcode.com/problems/count-of-smaller-numbers-after-self/
*/

package leetcode

func countSmaller(nums []int) []int {
	counts, clone := make([]int, len(nums)), make([]int, len(nums))
	indices1, indices2 := make([]int, len(nums)), make([]int, len(nums))
	for i := range indices1 {
		indices1[i] = i
	}

	helper315(nums, counts, clone, indices1, indices2, 0, len(nums)-1)
	return counts
}

func helper315(nums, counts, clone, indices1, indices2 []int, start, end int) {
	merge := func(nums, clone, indices1, indices2 []int, start, mid, end int) {
		i, j, k := start, mid+1, start
		for i <= mid || j <= end {
			if j > end || (i <= mid && nums[i] < nums[j]) {
				clone[k], indices2[k] = nums[i], indices1[i]
				i, k = i+1, k+1
			} else {
				clone[k], indices2[k] = nums[j], indices1[j]
				j, k = j+1, k+1
			}
		}

		for i = start; i <= end; i++ {
			nums[i], indices1[i] = clone[i], indices2[i]
		}
	}

	if start < end {
		mid := (start + end) / 2
		helper315(nums, counts, clone, indices1, indices2, start, mid)
		helper315(nums, counts, clone, indices1, indices2, mid+1, end)

		i, j := start, mid+1
		for ; i <= mid; i++ {
			for ; j <= end && nums[j] < nums[i]; j++ {
			}

			counts[indices1[i]] += j - mid - 1
		}

		merge(nums, clone, indices1, indices2, start, mid, end)
	}
}
