/*
LeetCode 444: https://leetcode.com/problems/sequence-reconstruction/
*/

package leetcode

func sequenceReconstruction(org []int, seqs [][]int) bool {
	if len(seqs) == 0 {
		return false
	}

	count := len(org)
	nexts, inDegrees := make(map[int][]int), make(map[int]int)
	for _, seq := range seqs {
		for i := 0; i < len(seq); i++ {
			if seq[i] < 1 || seq[i] > count {
				return false
			}

			if i > 0 {
				nexts[seq[i-1]] = append(nexts[seq[i-1]], seq[i])
				inDegrees[seq[i]]++
			}
		}
	}

	queue, built := make([]int, 0), make([]int, 0)
	for num := 1; num <= count; num++ {
		if inDegrees[num] == 0 {
			queue = append(queue, num)
		}
	}

	for len(queue) == 1 {
		num := queue[0]
		queue = queue[1:]
		built = append(built, num)
		for _, next := range nexts[num] {
			inDegrees[next]--
			if inDegrees[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	if len(org) != len(built) {
		return false
	}

	for i := range org {
		if org[i] != built[i] {
			return false
		}
	}

	return true
}
