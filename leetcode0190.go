/*
LeetCode 190: https://leetcode.com/problems/reverse-bits/
*/

package leetcode

func reverseBits(num uint32) uint32 {
	var result uint32
	for i := 0; i < 32; i++ {
		bit := num & 0x80000000
		num = num << 1
		result = (result >> 1) | bit
	}

	return result
}
