/*
LeetCode 888: https://leetcode.com/problems/fair-candy-swap/
*/

package leetcode

func fairCandySwap(A []int, B []int) []int {
	sum1 := 0
	nums1 := make(map[int]bool)
	for _, n := range A {
		sum1 += n
		nums1[n] = true
	}

	sum2 := 0
	nums2 := make(map[int]bool)
	for _, n := range B {
		sum2 += n
		nums2[n] = true
	}

	diff := (sum1 - sum2) / 2
	for n1 := range nums1 {
		if _, exists := nums2[n1-diff]; exists {
			return []int{n1, n1 - diff}
		}
	}

	return make([]int, 0)
}
