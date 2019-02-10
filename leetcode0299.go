/*
LeetCode 299: https://leetcode.com/problems/bulls-and-cows/
*/

package leetcode

import "fmt"

func getHint(secret string, guess string) string {
	digits := make([]int, 10)
	for _, digit := range secret {
		digits[int(digit-'0')]++
	}

	bull := 0
	for i := 0; i < len(guess); i++ {
		if secret[i] == guess[i] {
			bull++
			digits[int(guess[i]-'0')]--
		}
	}

	cow := 0
	for i := 0; i < len(guess); i++ {
		if secret[i] != guess[i] && digits[int(guess[i]-'0')] > 0 {
			cow++
			digits[int(guess[i]-'0')]--
		}
	}

	return fmt.Sprintf("%dA%dB", bull, cow)
}
