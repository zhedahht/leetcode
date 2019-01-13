/*
LeetCode 974: https://leetcode.com/problems/subarray-sums-divisible-by-k/
*/

package leetcode

func subarraysDivByK(A []int, K int) int {
	sums := make([]int, K)
	sums[0] = 1
	sum := 0
	count := 0
	for _, num := range A {
		sum += num
		index := sum % K
		if index < 0 {
			index += K
		}

		count += sums[index]
		sums[index]++
	}

	return count
}
