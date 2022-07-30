/*
LeetCode 465: https://leetcode.com/problems/optimal-account-balancing/
*/

package leetcode

import "math"

func minTransfers(transactions [][]int) int {
	vals := make(map[int]int)
	for _, t := range transactions {
		vals[t[0]] -= t[2]
		vals[t[1]] += t[2]
	}

	pos := make(map[int]int)
	neg := make(map[int]int)
	for _, val := range vals {
		if val > 0 {
			pos[val]++
		} else if val < 0 {
			neg[-val]++
		}
	}

	return helper0465(pos, neg, 0, len(transactions))
}

func helper0465(pos, neg map[int]int, count int, limit int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	add := func(nums map[int]int, val int) {
		nums[val]++
	}

	remove := func(nums map[int]int, val int) {
		nums[val]--
		if nums[val] == 0 {
			delete(nums, val)
		}
	}

	clone := func(nums map[int]int) map[int]int {
		result := make(map[int]int)
		for k, v := range nums {
			result[k] = v
		}
		return result
	}

	if len(pos) == 0 {
		return count
	}

	if min(len(pos), len(neg))+count >= limit {
		return limit
	}

	toBeRemoved := make(map[int]int)
	for p := range pos {
		if neg[p] > 0 {
			toBeRemoved[p] = min(pos[p], neg[p])
		}
	}

	if len(toBeRemoved) > 0 {
		removed := 0
		for k, v := range toBeRemoved {
			for i := 0; i < v; i++ {
				remove(pos, k)
				remove(neg, k)
			}
			removed += v
		}

		result := helper0465(pos, neg, count+removed, limit)

		for k, v := range toBeRemoved {
			for i := 0; i < v; i++ {
				add(pos, k)
				add(neg, k)
			}
		}

		return result
	}

	clonePos := clone(pos)
	cloneNeg := clone(neg)

	result := math.MaxInt
	for p := range clonePos {
		for n := range cloneNeg {
			remove(pos, p)
			remove(neg, n)
			if p > n {
				add(pos, p-n)
				steps := helper0465(pos, neg, count+1, min(result, limit))
				if result > steps {
					result = steps
				}
				remove(pos, p-n)
			} else {
				add(neg, n-p)
				steps := helper0465(pos, neg, count+1, min(result, limit))
				if result > steps {
					result = steps
				}
				remove(neg, n-p)
			}
			add(pos, p)
			add(neg, n)
		}
	}

	return result
}
