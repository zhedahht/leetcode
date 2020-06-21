/*
LeetCode 914: https://leetcode.com/problems/x-of-a-kind-in-a-deck-of-cards/
*/

package leetcode

func hasGroupsSizeX(deck []int) bool {
	valToCount := make(map[int]int)
	var count int
	for _, val := range deck {
		valToCount[val]++
		count = valToCount[val]
	}

	gcdVal := count
	for _, c := range valToCount {
		gcdVal = gcd(gcdVal, c)
	}

	return gcdVal >= 2
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}
