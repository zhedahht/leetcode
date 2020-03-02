/*
LeetCode 784: https://leetcode.com/problems/letter-case-permutation/
*/

package leetcode

func letterCasePermutation(S string) []string {
	return helper784(S, 0, make([]byte, 0), make([]string, 0))
}

func helper784(S string, i int, perm []byte, result []string) []string {
	if i == len(S) {
		result = append(result, string(perm))
		return result
	}

	perm = append(perm, S[i])
	result = helper784(S, i+1, perm, result)
	perm = perm[:len(perm)-1]
	if S[i] >= 'A' && S[i] <= 'Z' {
		perm = append(perm, S[i]+32)
		result = helper784(S, i+1, perm, result)
		perm = perm[:len(perm)-1]
	} else if S[i] >= 'a' && S[i] <= 'z' {
		perm = append(perm, S[i]-32)
		result = helper784(S, i+1, perm, result)
		perm = perm[:len(perm)-1]
	}

	return result
}
