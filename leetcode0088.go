/*
LeetCode 88: https://leetcode.com/problems/merge-sorted-array/
*/

package leetcode

// NOTE: the name should merge. Rename it to avoid conflicts with others.
func merge88(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for i >= 0 || j >= 0 {
		if j < 0 || (i >= 0 && nums1[i] >= nums2[j]) {
			nums1[k] = nums1[i]
			k, i = k-1, i-1
		} else {
			nums1[k] = nums2[j]
			k, j = k-1, j-1
		}
	}
}
