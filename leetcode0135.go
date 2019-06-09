/*
LeetCode 135: https://leetcode.com/problems/candy/
*/

package leetcode

func candy(ratings []int) int {
	queue1, queue2, counts := make([]int, 0), make([]int, 0), make([]int, len(ratings))
	for i, r := range ratings {
		if (i == 0 || r <= ratings[i-1]) && (i == len(ratings)-1 || r <= ratings[i+1]) {
			queue1 = append(queue1, i)
		}
	}

	count := 1
	for len(queue1) > 0 {
		i := queue1[0]
		queue1 = queue1[1:]
		counts[i] = count

		if i > 0 && ratings[i-1] > ratings[i] && counts[i-1] <= counts[i] {
			queue2 = append(queue2, i-1)
		}

		if i < len(ratings)-1 && ratings[i+1] > ratings[i] && counts[i+1] <= counts[i] {
			queue2 = append(queue2, i+1)
		}

		if len(queue1) == 0 {
			count++
			queue1, queue2 = queue2, make([]int, 0)
		}
	}

	result := 0
	for _, c := range counts {
		result += c
	}

	return result
}
