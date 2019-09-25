/*
LeetCode 179: https://leetcode.com/problems/largest-number/
*/

package leetcode

import (
	"fmt"
	"sort"
	"strings"
)

func largestNumber(nums []int) string {
	sort.Slice(nums, func(m, n int) bool {
		mn := fmt.Sprintf("%d%d", nums[m], nums[n])
		nm := fmt.Sprintf("%d%d", nums[n], nums[m])
		return strings.Compare(mn, nm) > 0
	})

	var result string
	for _, num := range nums {
		if result == "0" {
			result = fmt.Sprintf("%d", num)
		} else {
			result = fmt.Sprintf("%s%d", result, num)
		}
	}

	return result
}
