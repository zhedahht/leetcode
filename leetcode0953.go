/*
LeetCode 953: https://leetcode.com/problems/verifying-an-alien-dictionary/
*/

package leetcode

func isAlienSorted(words []string, order string) bool {
	isLess := func(str1, str2 string, dict []int) bool {
		for i := 0; i < len(str1) && i < len(str2); i++ {
			idx1 := dict[int(str1[i]-'a')]
			idx2 := dict[int(str2[i]-'a')]
			if idx1 < idx2 {
				return true
			}

			if idx1 > idx2 {
				return false
			}
		}

		return len(str1) <= len(str2)
	}

	dict := make([]int, 26)
	for i, b := range order {
		dict[int(b-'a')] = i
	}

	for i := 0; i < len(words)-1; i++ {
		if !isLess(words[i], words[i+1], dict) {
			return false
		}
	}

	return true
}
