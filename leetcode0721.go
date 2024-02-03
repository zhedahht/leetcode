/*
LeetCode 721: https://leetcode.com/problems/accounts-merge/description/
*/

package leetcode

import "slices"

func accountsMerge(accounts [][]string) [][]string {
	emailToParent := make(map[string]string)
	emailToName := make(map[string]string)
	for _, account := range accounts {
		name := account[0]
		for i := 1; i < len(account); i++ {
			email := account[i]
			emailToName[email] = name
			union721(email, account[1], emailToParent)
		}
	}

	parentToEmails := make(map[string][]string)
	for email := range emailToName {
		parent := findParent721(email, emailToParent)
		if _, exists := parentToEmails[parent]; !exists {
			parentToEmails[parent] = make([]string, 0)
		}

		parentToEmails[parent] = append(parentToEmails[parent], email)
	}

	result := make([][]string, 0)
	for parent, emails := range parentToEmails {
		account := []string{emailToName[parent]}
		slices.Sort(emails)
		account = append(account, emails...)
		result = append(result, account)
	}

	return result
}

func union721(email1, email2 string, emailToParent map[string]string) {
	if _, exists := emailToParent[email1]; !exists {
		emailToParent[email1] = email1
	}

	parent1 := findParent721(email1, emailToParent)
	parent2 := findParent721(email2, emailToParent)
	emailToParent[parent1] = parent2
}

func findParent721(email string, emailToParent map[string]string) string {
	parent := emailToParent[email]
	if parent != emailToParent[parent] {
		parent = findParent721(parent, emailToParent)
		emailToParent[email] = parent
	}

	return parent
}
