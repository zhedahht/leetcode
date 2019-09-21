/*
LeetCode 187: https://leetcode.com/problems/repeated-dna-sequences/
*/

package leetcode

func findRepeatedDnaSequences(s string) []string {
	result := make([]string, 0)
	if len(s) < 10 {
		return result
	}

	dict := make(map[string]int)
	for i := 0; i < len(s)-9; i++ {
		sub := string(s[i : i+10])
		if _, exists := dict[sub]; !exists {
			dict[sub] = 1
		} else {
			dict[sub]++
		}
	}

	for k, v := range dict {
		if v > 1 {
			result = append(result, k)
		}
	}

	return result
}
