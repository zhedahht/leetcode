/*
LeetCode 442: https://leetcode.com/problems/find-all-duplicates-in-an-array/description/
*/

package leetcode

func findDuplicates(nums []int) []int {
	dups := make([]int, 0)
	for i, num := range nums {
		idx, n := i, num
		if n != idx+1 {
			nums[idx] = 0
			for n > 0 && n != idx+1 {
				next := nums[n-1]
				if n == next {
					dups = append(dups, n)
					break
				} else {
					nums[n-1] = n
					idx = n - 1
					n = next
				}
			}
		}
	}

	return dups
}
