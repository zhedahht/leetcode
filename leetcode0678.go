/*
LeetCode 678: https://leetcode.com/problems/valid-parenthesis-string/description/
*/

package leetcode

func checkValidString(s string) bool {
	left, right, star := []int{}, []int{}, []int{}
	for i, r := range s {
		switch r {
		case '(':
			left = append(left, i)
		case ')':
			if len(left) > 0 {
				left = left[:len(left)-1]
			} else {
				right = append(right, i)
			}
		default:
			star = append(star, i)
		}
	}

	for len(left) > 0 && len(star) > 0 {
		if left[len(left)-1] > star[len(star)-1] {
			return false
		}

		left = left[:len(left)-1]
		star = star[:len(star)-1]
	}

	for len(right) > 0 && len(star) > 0 {
		if right[0] < star[0] {
			return false
		}

		right = right[1:]
		star = star[1:]
	}

	return len(left) == 0 && len(right) == 0
}
