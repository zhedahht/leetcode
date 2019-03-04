/*
LeetCode 990: https://leetcode.com/problems/satisfiability-of-equality-equations/
*/

package leetcode

func equationsPossible(equations []string) bool {
	chToParent := make([]byte, 26)
	for _, equation := range equations {
		ch1 := equation[0]
		ch2 := equation[3]
		chToParent[ch1-'a'] = ch1
		chToParent[ch2-'a'] = ch2
	}

	for _, equation := range equations {
		if equation[1] == '=' {
			ch1 := equation[0]
			ch2 := equation[3]
			union(chToParent, ch1, ch2)
		}
	}

	for _, equation := range equations {
		if equation[1] == '!' {
			ch1 := equation[0]
			ch2 := equation[3]

			if getParent(chToParent, ch1) == getParent(chToParent, ch2) {
				return false
			}
		}
	}

	return true
}

func getParent(chToParent []byte, ch byte) byte {
	parent := chToParent[ch-'a']
	if ch != parent {
		parent = getParent(chToParent, parent)
		chToParent[ch-'a'] = parent
	}

	return parent
}

func union(chToParent []byte, ch1, ch2 byte) {
	parent1 := getParent(chToParent, ch1)
	parent2 := getParent(chToParent, ch2)
	chToParent[parent1-'a'] = parent2
}
