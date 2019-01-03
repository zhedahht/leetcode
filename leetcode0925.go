/*
LeetCode 925: https://leetcode.com/problems/long-pressed-name/
*/

package leetcode

func isLongPressedName(name string, typed string) bool {
	i := 0
	j := 0
	for i < len(name) && j < len(typed) {
		var ch1, ch2 byte
		var count1, count2 int
		ch1, i, count1 = parseName(name, i)
		ch2, j, count2 = parseName(typed, j)
		if ch1 != ch2 || count1 > count2 {
			return false
		}
	}

	return i == len(name) && j == len(typed)
}

func parseName(str string, index int) (byte, int, int) {
	ch := str[index]
	count := 0
	for ; index < len(str) && str[index] == ch; index++ {
		count++
	}

	return ch, index, count
}
