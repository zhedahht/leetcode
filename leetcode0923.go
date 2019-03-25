/*
LeetCode 923: https://leetcode.com/problems/3sum-with-multiplicity/
*/

package leetcode

import "sort"

func threeSumMulti(A []int, target int) int {
	sort.Ints(A)
	count := 0
	for i := 0; i < len(A); i++ {
		j, k := i+1, len(A)-1
		for j < k {
			if A[i]+A[j]+A[k] > target {
				k--
			} else if A[i]+A[j]+A[k] < target {
				j++
			} else if A[j] != A[k] {
				sameIndex1 := j
				for sameIndex1 < k && A[sameIndex1] == A[j] {
					sameIndex1++
				}

				sameIndex2 := k
				for sameIndex2 > j && A[sameIndex2] == A[k] {
					sameIndex2--
				}

				count = (count + (sameIndex1-j)*(k-sameIndex2)) % 1000000007
				j, k = sameIndex1, sameIndex2
			} else {
				length := k - j + 1
				count = (count + length*(length-1)/2) % 1000000007
				j = k + 1
			}
		}
	}

	return count
}
