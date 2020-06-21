/*
LeetCode 991: https://leetcode.com/problems/broken-calculator/
*/

package leetcode

func brokenCalc(X int, Y int) int {
	if X >= Y {
		return X - Y
	}

	if Y%2 == 0 {
		return 1 + brokenCalc(X, Y/2)
	}

	return 1 + brokenCalc(X, Y+1)
}
