/*
LeetCode 972: https://leetcode.com/problems/equal-rational-numbers/
*/

package leetcode

import "bytes"

func isRationalEqual(S string, T string) bool {
	int1, nonRepeating1, repeating1 := normalizeRational(S)
	int2, nonRepeating2, repeating2 := normalizeRational(T)
	return int1 == int2 && nonRepeating1 == nonRepeating2 && repeating1 == repeating2
}

func normalizeRational(num string) (string, string, string) {
	i := 0
	intArray := make([]byte, 0)
	for i < len(num) && num[i] != '.' {
		intArray = append(intArray, num[i])
		i++
	}

	if i == len(num) {
		return string(intArray), "", ""
	}

	i++
	nonRepeatingArray := make([]byte, 0)
	for i < len(num) && num[i] != '(' {
		nonRepeatingArray = append(nonRepeatingArray, num[i])
		i++
	}

	if i == len(num) {
		return string(intArray), string(nonRepeatingArray), ""
	}

	repeatingArray := []byte(num)[i+1 : len(num)-1]

	repeatingArray = shortenRepeating(repeatingArray)
	if len(repeatingArray) == 1 && repeatingArray[0] == '0' {
		repeatingArray = make([]byte, 0)
	}

	if len(repeatingArray) == 0 {
		k := 0
		for ; k < len(nonRepeatingArray); k++ {
			if nonRepeatingArray[0] != '0' {
				break
			}
		}

		if k == len(nonRepeatingArray) {
			nonRepeatingArray = make([]byte, 0)
		}
	}

	nonRepeatingArray, repeatingArray = mergeRepeating(nonRepeatingArray, repeatingArray)
	if len(repeatingArray) == 1 && repeatingArray[0] == '9' {
		repeatingArray = []byte{}
		intArray, nonRepeatingArray = addOne(intArray, nonRepeatingArray)
	}

	return string(intArray), string(nonRepeatingArray), string(repeatingArray)
}

func shortenRepeating(repeatingArray []byte) []byte {
	if len(repeatingArray) == 1 {
		return repeatingArray
	}

	for i := 1; i < len(repeatingArray); i++ {
		if len(repeatingArray)%i == 0 {
			j := 0
			seg := repeatingArray[:i]
			for j = i; j < len(repeatingArray); j += i {
				if !bytes.Equal(seg, repeatingArray[j:j+i]) {
					break
				}

				seg = repeatingArray[j : j+i]
			}

			if j == len(repeatingArray) {
				return seg
			}
		}
	}

	return repeatingArray
}

func mergeRepeating(nonRepeatingArray, repeatingArray []byte) ([]byte, []byte) {
	if len(nonRepeatingArray) == 0 || len(repeatingArray) == 0 {
		return nonRepeatingArray, repeatingArray
	}

	window := make([]byte, len(repeatingArray))
	copy(window, repeatingArray)

	i, j := len(nonRepeatingArray)-1, len(repeatingArray)-2
	last := window[len(window)-1]
	for ; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if nonRepeatingArray[i] != last {
			break
		}

		window = make([]byte, 0)
		window = append(window, nonRepeatingArray[i:]...)
		window = append(window, repeatingArray[:j+1]...)
		last = window[len(window)-1]
	}

	for ; j == -1 && i >= 0; i-- {
		if nonRepeatingArray[i] != last {
			break
		}

		window = nonRepeatingArray[i : i+len(repeatingArray)]
		last = window[len(window)-1]
	}

	return nonRepeatingArray[:i+1], window
}

func addOne(intArray, nonRepeatingArray []byte) ([]byte, []byte) {
	if len(nonRepeatingArray) == 0 {
		intArray = addOneInternal(intArray)
	} else {
		nonRepeatingArray = addOneInternal(nonRepeatingArray)
	}

	return intArray, nonRepeatingArray
}

func addOneInternal(num []byte) []byte {
	val := 1
	result := make([]byte, 0)
	for i := len(num) - 1; i >= 0; i-- {
		digit := int(num[i]-'0') + val
		if digit >= 10 {
			digit -= 10
			val = 1
		} else {
			val = 0
		}

		result = append(result, byte(digit)+'0')
	}

	if val == 1 {
		result = append(result, '1')
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}
