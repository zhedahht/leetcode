/*
LeetCode 413: https://leetcode.com/problems/arithmetic-slices/description/
*/

package leetcode

func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	subs := make([]int, 0)
	length := 2
	diff := nums[1] - nums[0]
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == diff {
			length++
		} else {
			if length >= 3 {
				subs = append(subs, length)
			}

			length = 2
			diff = nums[i] - nums[i-1]
		}
	}

	if length >= 3 {
		subs = append(subs, length)
	}

	result := 0
	lengthToCount := make(map[int]int)
	for _, l := range subs {
		count, exists := lengthToCount[l]
		if exists {
			result += count
		} else {
			n := l - 2
			count = n * (1 + n) / 2
			result += count
			lengthToCount[l] = count
		}
	}

	return result
}
