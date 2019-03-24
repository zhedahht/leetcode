/*
LeetCode 984: https://leetcode.com/problems/string-without-aaa-or-bbb/
*/

package leetcode

func strWithout3a3b(A int, B int) string {
	big, small := A, B
	bigCh, smallCh := byte(97), byte(98)
	if A < B {
		big, small = B, A
		bigCh, smallCh = byte(98), byte(97)
	}

	result := make([]byte, 0)
	for big > small && big > 1 && small > 0 {
		result = append(result, bigCh, bigCh, smallCh)
		big, small = big-2, small-1
	}

	for big > 0 && small > 0 {
		result = append(result, bigCh, smallCh)
		big, small = big-1, small-1
	}

	for big > 0 {
		result = append(result, bigCh)
		big--
	}

	return string(result)
}
