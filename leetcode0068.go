/*
LeetCode 68: https://leetcode.com/problems/text-justification/
*/

package leetcode

func fullJustify(words []string, maxWidth int) []string {
	result := make([]string, 0)
	for i := 0; i < len(words); {
		j, total := i, len(words[i])
		for j < len(words)-1 && total+len(words[j+1]) < maxWidth {
			j, total = j+1, total+len(words[j+1])+1
		}

		diff, cur := maxWidth-total, words[i]
		if j > i && j < len(words)-1 {
			count, average := diff%(j-i), diff/(j-i)
			for k := i + 1; k <= j; k++ {
				for n := 0; n <= average; n++ {
					cur += " "
				}

				if k-i <= count {
					cur += " "
				}

				cur += words[k]
			}
		} else {
			for k := i + 1; k <= j; k++ {
				cur += " " + words[k]
			}

			for k := 0; k < diff; k++ {
				cur += " "
			}
		}

		result, i = append(result, cur), j+1
	}

	return result
}
