/*
LeetCode 438: https://leetcode.com/problems/find-all-anagrams-in-a-string/description/
*/

package leetcode

func findAnagrams(s string, p string) []int {
	result := make([]int, 0)
	if len(s) < len(p) {
		return result
	}

	countP := make(map[byte]int)
	for _, b := range []byte(p) {
		countP[b]++
	}

	bytesS := []byte(s)
	countS := make(map[byte]int)
	for _, b := range bytesS[:len(p)-1] {
		countS[b]++
	}

	for i := len(p) - 1; i < len(s); i++ {
		countS[bytesS[i]]++
		if len(countS) == len(countP) {
			same := true
			for b, c := range countP {
				if countS[b] != c {
					same = false
					break
				}
			}

			if same {
				result = append(result, i-len(p)+1)
			}
		}

		first := bytesS[i-len(p)+1]
		countS[first]--
		if countS[first] == 0 {
			delete(countS, first)
		}
	}

	return result
}
