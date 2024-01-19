/*
LeetCode 657: https://leetcode.com/problems/robot-return-to-origin/description/
*/

package leetcode

func judgeCircle(moves string) bool {
	counts := make(map[rune]int)
	for _, r := range moves {
		counts[r]++
	}

	return counts[rune('U')] == counts[rune('D')] && counts[rune('L')] == counts[rune('R')]
}
