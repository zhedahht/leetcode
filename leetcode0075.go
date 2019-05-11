/*
LeetCode 75: https://leetcode.com/problems/sort-colors/
*/

package leetcode

func sortColors(nums []int) {
	i, j := -1, len(nums)
	for k := 0; k < j; {
		if nums[k] == 0 && k != i {
			i++
			nums[i], nums[k] = nums[k], nums[i]
		} else if nums[k] == 2 {
			j--
			nums[j], nums[k] = nums[k], nums[j]
		} else {
			k++
		}
	}
}
