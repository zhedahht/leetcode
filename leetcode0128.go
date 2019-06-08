/*
LeetCode 128: https://leetcode.com/problems/longest-consecutive-sequence/submissions/
*/

package leetcode

func longestConsecutive(nums []int) int {
	parentMap, countMap := make(map[int]int), make(map[int]int)
	for _, num := range nums {
		if _, exists := parentMap[num]; exists {
			continue
		}

		parentMap[num], countMap[num] = num, 1
		neighbors := [2]int{num - 1, num + 1}
		for _, n := range neighbors {
			if _, exists := parentMap[n]; exists {
				count1 := countMap[getParent128(n, parentMap)]
				count2 := countMap[getParent128(num, parentMap)]

				union128(n, num, parentMap)
				countMap[getParent128(num, parentMap)] = count1 + count2
			}
		}
	}

	max := 0
	for _, count := range countMap {
		if count > max {
			max = count
		}
	}

	return max
}

func getParent128(num int, parentMap map[int]int) int {
	p := parentMap[num]
	if p != parentMap[p] {
		p = getParent128(p, parentMap)
		parentMap[num] = p
	}

	return p
}

func union128(num1, num2 int, parentMap map[int]int) {
	p1, p2 := getParent128(num1, parentMap), getParent128(num2, parentMap)
	parentMap[p1] = p2
}
