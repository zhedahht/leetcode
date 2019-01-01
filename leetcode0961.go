/*
LeetCode 961: https://leetcode.com/problems/n-repeated-element-in-size-2n-array/
*/

package leetcode

func repeatedNTimes(A []int) int {
	queue := make([]int, 0)
	for _, num := range A {
		for _, n := range queue {
			if n == num {
				return num
			}
		}

		queue = append(queue, num)
		if len(queue) > 4 {
			queue = queue[1:]
		}
	}

	return -1
}
