/*
LeetCode 43: https://leetcode.com/problems/multiply-strings/
*/

package leetcode

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	digitResult := make(map[byte][]byte)
	if len(num1) > len(num2) {
		return string(helper43([]byte(num2), []byte(num1), 0, len(num1)-1, digitResult))
	}

	return string(helper43([]byte(num1), []byte(num2), 0, len(num2)-1, digitResult))
}

func helper43(num1, num2 []byte, left, right int, digitResult map[byte][]byte) []byte {
	result := make([]byte, 0)

	if left == right {
		var temp []byte
		if r, exists := digitResult[num2[left]]; exists {
			temp = r
		} else {
			temp = multiplyDigit(num1, num2[left])
			digitResult[num2[left]] = temp
		}

		result := make([]byte, len(temp))
		copy(result, temp)
		for i := left + 1; i < len(num2); i++ {
			result = append(result, '0')
		}

		return result
	}

	mid := (left + right) / 2
	result1 := helper43(num1, num2, left, mid, digitResult)
	result2 := helper43(num1, num2, mid+1, right, digitResult)

	i, j, carry := len(result1)-1, len(result2)-1, byte(0)
	for i >= 0 || j >= 0 {
		digit1, digit2 := byte('0'), byte('0')
		if i >= 0 {
			digit1 = result1[i]
			i--
		}

		if j >= 0 {
			digit2 = result2[j]
			j--
		}

		sum := (digit1 - '0') + (digit2 - '0') + carry
		if sum >= 10 {
			sum -= byte(10)
			carry = byte(1)
		} else {
			carry = byte(0)
		}

		result = append(result, sum+'0')
	}

	if carry > 0 {
		result = append(result, '1')
	}

	reverseBytes(result)
	return result
}

func reverseBytes(bytes []byte) {
	i, j := 0, len(bytes)-1
	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i, j = i+1, j-1
	}
}

func multiplyDigit(num []byte, b byte) []byte {
	result := make([]byte, 0)
	carry := byte(0)
	for i := len(num) - 1; i >= 0; i-- {
		val := (b-'0')*(num[i]-'0') + carry
		carry = val / 10
		val = val % 10
		result = append(result, val+'0')
	}

	if carry > 0 {
		result = append(result, carry+'0')
	}

	reverseBytes(result)
	return result
}
