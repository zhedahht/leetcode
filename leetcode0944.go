/*
LeetCode 944: https://leetcode.com/problems/delete-columns-to-make-sorted/
*/

package leetcode

func minDeletionSize(A []string) int {
	count := 0
	for i := range A[0] {
		prev := byte('a')
		for _, str := range A {
			if str[i] < prev {
				count++
				break
			}

			prev = str[i]
		}
	}

	return count
}
