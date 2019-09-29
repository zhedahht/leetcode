/*
LeetCode 168: https://leetcode.com/problems/excel-sheet-column-title/
*/

package leetcode

func convertToTitle(n int) string {
	digits := make([]byte, 0)
	for n > 0 {
		d := (n-1)%26 + 1
		n = (n - 1) / 26
		digits = append(digits, byte(d-1)+'A')
	}

	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	return string(digits)
}
