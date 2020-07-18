/*
LeetCode 752: https://leetcode.com/problems/open-the-lock/
*/

package leetcode

func openLock(deadends []string, target string) int {
	getNext := func(s string) []string {
		next := make([]string, 0)
		for i := 0; i < len(s); i++ {
			b1, b2 := []byte(s), []byte(s)
			if b1[i] == byte('9') {
				b1[i] = byte('0')
			} else {
				b1[i]++
			}

			next = append(next, string(b1))

			if b2[i] == byte('0') {
				b2[i] = byte('9')
			} else {
				b2[i]--
			}

			next = append(next, string(b2))
		}

		return next
	}

	dead := make(map[string]bool)
	for _, d := range deadends {
		dead[d] = true
	}

	visited := make(map[string]bool)
	queue1, queue2 := make([]string, 0), make([]string, 0)
	const init = "0000"
	if !dead[init] {
		queue1 = append(queue1, init)
		visited[init] = true
	}

	steps := 0
	for len(queue1) > 0 {
		s := queue1[0]
		queue1 = queue1[1:]
		if s == target {
			return steps
		}

		next := getNext(s)
		for _, n := range next {
			if !dead[n] && !visited[n] {
				queue2 = append(queue2, n)
				visited[n] = true
			}
		}

		if len(queue1) == 0 {
			queue1 = queue2
			queue2 = make([]string, 0)
			steps++
		}
	}

	return -1
}
