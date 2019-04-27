/*
LeetCode 15: https://leetcode.com/problems/3sum/
*/

package leetcode

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; i < len(nums)-2; {
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum < 0 {
				j++
			} else if sum > 0 {
				k--
			} else {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				temp := j
				for j < k && nums[j] == nums[temp] {
					j++
				}
			}
		}

		temp := i
		for i < len(nums)-2 && nums[i] == nums[temp] {
			i++
		}
	}

	return result
}
