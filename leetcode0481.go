/*
LeetCode 481: https://leetcode.com/problems/magical-string/description/
*/

package leetcode

func magicalString(n int) int {
	digits := []int{1, 2, 2}
	slow, fast := 1, 2
	for len(digits) <= n {
		count := digits[slow+1]
		digit := 3 - digits[fast]

		for i := 0; i < count; i++ {
			digits = append(digits, digit)
		}

		fast += count
		slow++
	}

	result := 0
	for i := 0; i < n; i++ {
		if digits[i] == 1 {
			result++
		}
	}

	return result
}
