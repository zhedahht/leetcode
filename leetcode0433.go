/*
LeetCode 433: https://leetcode.com/problems/minimum-genetic-mutation/description/
*/

package leetcode

func minMutation(startGene string, endGene string, bank []string) int {
	all := make(map[string]bool)
	for _, gene := range bank {
		all[gene] = true
	}

	if !all[endGene] {
		return -1
	}

	queue1, queue2 := make([]string, 0), make([]string, 0)
	visited := make(map[string]bool)
	queue1 = append(queue1, startGene)
	visited[startGene] = true
	steps := 0
	found := false
	for len(queue1) > 0 {
		gene := queue1[0]
		queue1 = queue1[1:]
		if gene == endGene {
			found = true
			break
		}

		for i := range gene {
			chars := []byte{'A', 'C', 'G', 'T'}
			for _, ch := range chars {
				if gene[i] != ch {
					nextBytes := []byte(gene)
					nextBytes[i] = ch
					next := string(nextBytes)
					if all[next] && !visited[next] {
						queue2 = append(queue2, next)
						visited[next] = true
					}
				}
			}
		}

		if len(queue1) == 0 {
			queue1 = queue2
			queue2 = make([]string, 0)
			steps++
		}
	}

	if !found {
		steps = -1
	}

	return steps
}
