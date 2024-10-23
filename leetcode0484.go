/*
LeetCode 484: https://leetcode.com/problems/find-permutation/description/
*/

package leetcode

func findPermutation(s string) []int {
	segs := make([][]int, 0)
	prev := byte(' ')
	for i := 0; i < len(s); i++ {
		if s[i] != prev {
			prev = s[i]
			if i == 0 {
				segs = append(segs, []int{1})
			} else {
				last := segs[len(segs)-1]
				lastLen := len(last)
				segs = append(segs, []int{last[lastLen-1]})
				segs[len(segs)-2] = last[:lastLen-1]
			}
		}

		if s[i] == 'D' {
			seg := []int{i + 2}
			seg = append(seg, segs[len(segs)-1]...)
			segs[len(segs)-1] = seg
		} else {
			segs[len(segs)-1] = append(segs[len(segs)-1], i+2)
		}
	}

	result := make([]int, 0)
	for _, seg := range segs {
		result = append(result, seg...)
	}

	return result
}
