/*
LeetCode 496: https://leetcode.com/problems/next-greater-element-i/description/
*/

package leetcode

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	numToPos := make(map[int]int)
	for i, num := range nums2 {
		numToPos[num] = i
	}

	result := make([]int, len(nums1))
	for i, num := range nums1 {
		pos := numToPos[num]
		j := pos + 1
		result[i] = -1
		for ; j < len(nums2); j++ {
			if nums2[j] > num {
				result[i] = nums2[j]
				break
			}
		}
	}

	return result
}
