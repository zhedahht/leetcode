/*
LeetCode 165: https://leetcode.com/problems/compare-version-numbers/
*/

package leetcode

func compareVersion(version1 string, version2 string) int {
	i1 := 0
	i2 := 0
	val1 := 0
	val2 := 0

	for i1 < len(version1) || i2 < len(version2) {
		val1, i1 = parseSeg(version1, i1)
		val2, i2 = parseSeg(version2, i2)

		if val1 > val2 {
			return 1
		}

		if val1 < val2 {
			return -1
		}
	}

	if i1 < len(version1) {
		return 1
	}

	if i2 < len(version2) {
		return -1
	}

	return 0
}

func parseSeg(version string, index int) (int, int) {
	val := 0
	for index < len(version) && version[index] != '.' {
		val = val*10 + int(version[index]-'0')
		index++
	}

	if index < len(version) {
		index++
	}

	return val, index
}
