/*
LeetCode 800: https://leetcode.com/problems/similar-rgb-color/
*/

package leetcode

func similarRGB(color string) string {
	result := make([]byte, 7)
	result[0] = color[0]
	for i := 0; i < 3; i++ {
		b := helper800(color[i*2+1], color[i*2+2])
		result[i*2+1], result[i*2+2] = b, b
	}

	return string(result)
}

func helper800(b1, b2 byte) byte {
	if b1 == b2 {
		return b1
	}

	b1, b2 = getColorValue(b1), getColorValue(b2)
	if b1 < b2 {
		delta1 := b2 - b1
		delta2 := 17 + b1 - b2
		if delta1 < delta2 {
			return toColorChar(b1)
		}

		return toColorChar(b1 + 1)
	}

	delta1 := b1 - b2
	delta2 := 17 + b2 - b1
	if delta1 < delta2 {
		return toColorChar(b1)
	}

	return toColorChar(b1 - 1)
}

func getColorValue(b byte) byte {
	if b >= '0' && b <= '9' {
		return byte(b - '0')
	}

	return byte(10 + b - 'a')
}

func toColorChar(b byte) byte {
	b = (b + 16) % 16
	if b <= 9 {
		return byte(b + '0')
	}

	return b - 10 + 'a'
}
