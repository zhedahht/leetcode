/*
LeetCode 415: https://leetcode.com/problems/add-strings/
*/

package leetcode

func addStrings(num1 string, num2 string) string {
	chArray := make([]byte, 0)
	carry := byte(0)
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		sum := carry
		if i >= 0 {
			sum += num1[i] - '0'
		}

		if j >= 0 {
			sum += num2[j] - '0'
		}

		if sum >= 10 {
			sum -= 10
			carry = 1
		} else {
			carry = 0
		}

		chArray = append(chArray, sum+'0')
	}

	if carry == 1 {
		chArray = append(chArray, '1')
	}

	for i := 0; i < len(chArray)/2; i++ {
		chArray[i], chArray[len(chArray)-1-i] = chArray[len(chArray)-1-i], chArray[i]
	}

	return string(chArray)
}
