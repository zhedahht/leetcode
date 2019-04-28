/*
LeetCode 18: https://leetcode.com/problems/4sum/
*/

package leetcode

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; i < len(nums)-3; {
		for j := i + 1; j < len(nums)-2; {
			k, m := j+1, len(nums)-1
			for k < m {
				sum := nums[i] + nums[j] + nums[k] + nums[m]
				if sum < target {
					k++
				} else if sum > target {
					m--
				} else {
					result = append(result, []int{nums[i], nums[j], nums[k], nums[m]})
					k = increaseIndex18(nums, k)
				}
			}

			j = increaseIndex18(nums, j)
		}

		i = increaseIndex18(nums, i)
	}

	return result
}

func increaseIndex18(nums []int, i int) int {
	temp := i
	for i < len(nums)-1 && nums[i] == nums[temp] {
		i++
	}

	return i
}
