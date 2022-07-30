/*
LeetCode 287: https://leetcode.com/problems/find-the-duplicate-number/
*/

package leetcode

func findDuplicate(nums []int) int {
	countRange := func(nums []int, left, right int) int {
		count := 0
		for _, num := range nums {
			if num >= left && num <= right {
				count++
			}
		}
		return count
	}

	left, right := 1, len(nums)-1
	for left <= right {
		if left == right {
			return left
		}
		mid := (left + right) / 2
		count := countRange(nums, left, mid)
		if count <= mid-left+1 {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return 0
}
