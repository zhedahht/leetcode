/*
LeetCode 327: https://leetcode.com/problems/count-of-range-sum/
*/

package leetcode

func countRangeSum(nums []int, lower int, upper int) int {
	sums := make([]int, len(nums)+1)
	for i, num := range nums {
		sums[i+1] = sums[i] + num
	}

	clone := make([]int, len(sums))
	return helper327(sums, clone, 0, len(sums)-1, lower, upper)
}

func helper327(sums, clone []int, start, end, lower, upper int) int {
	if start >= end {
		return 0
	}

	mid := (start + end) / 2
	count := helper327(sums, clone, start, mid, lower, upper)
	count += helper327(sums, clone, mid+1, end, lower, upper)
	i, j, k := start, mid+1, mid+1
	for ; i <= mid; i++ {
		for j <= end && sums[j] < sums[i]+lower {
			j++
		}

		for k <= end && sums[k] <= sums[i]+upper {
			k++
		}

		count += k - j
	}

	for i, j, k = start, mid+1, start; i <= mid || j <= end; {
		if i == mid+1 || (j <= end && sums[i] > sums[j]) {
			clone[k] = sums[j]
			k, j = k+1, j+1
		} else {
			clone[k] = sums[i]
			k, i = k+1, i+1
		}
	}

	for i = start; i <= end; i++ {
		sums[i] = clone[i]
	}

	return count
}
