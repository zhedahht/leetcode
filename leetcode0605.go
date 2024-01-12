/*
LeetCode 605: https://leetcode.com/problems/can-place-flowers/description/
*/

package leetcode

func canPlaceFlowers(flowerbed []int, n int) bool {
	prev1 := false
	count := 0
	for i, val := range flowerbed {
		if val == 0 {
			if !prev1 && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
				count++
				prev1 = true
			} else {
				prev1 = false
			}
		} else {
			prev1 = true
		}
	}

	return count >= n
}
