/*
https://leetcode.com/problems/subarray-sum-equals-k/description/
*/

package leetcode

func subarraySum(nums []int, k int) int {
	sumCounts := make(map[int]int)
	sumCounts[0] = 1
	result := 0
	sum := 0
	for _, num := range nums {
		sum += num
		result += sumCounts[sum-k]
		sumCounts[sum]++
	}

	return result
}
