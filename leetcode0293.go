/*
LeetCode 293: https://leetcode.com/problems/flip-game/
*/

package leetcode

func generatePossibleNextMoves(s string) []string {
	result := make([]string, 0)
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '+' && s[i+1] == '+' {
			chars := []byte(s)
			chars[i], chars[i+1] = '-', '-'
			result = append(result, string(chars))
		}
	}

	return result
}
