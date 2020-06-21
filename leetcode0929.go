/*
LeetCode 929: https://leetcode.com/problems/unique-email-addresses/
*/

package leetcode

func numUniqueEmails(emails []string) int {
	set := make(map[string]bool)
	for _, email := range emails {
		normalized := normalizeEmail(email)
		if _, ok := set[normalized]; !ok {
			set[normalized] = true
		}
	}

	return len(set)
}

func normalizeEmail(email string) string {
	chArray := make([]byte, 0)
	i := 0
	for ; i < len(email) && email[i] != '@'; i++ {
		ch := email[i]
		if ch == '.' {
			continue
		} else if ch == '+' {
			break
		} else {
			chArray = append(chArray, ch)
		}
	}

	for ; i < len(email); i++ {
		if email[i] == '@' {
			break
		}
	}

	for ; i < len(email); i++ {
		chArray = append(chArray, email[i])
	}

	return string(chArray)
}
