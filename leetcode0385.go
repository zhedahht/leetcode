/*
LeetCode 385: https://leetcode.com/problems/mini-parser/
*/

package leetcode

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */
func deserialize(s string) *NestedInteger {
	_, result := helper([]byte(s), 0)
	return result
}

func helper(s []byte, index int) (int, *NestedInteger) {
	if s[index] == '[' {
		result := &NestedInteger{}
		index, result = parseList(s, index+1, result)
		index++ //for ']'

		return index, result
	}

	return parseInteger(s, index)
}

func parseList(s []byte, index int, parent *NestedInteger) (int, *NestedInteger) {
	for index < len(s) && s[index] != ']' {
		if s[index] == '[' {
			var nested *NestedInteger
			index, nested = helper(s, index)
			parent.Add(*nested)
		} else {
			var nested *NestedInteger
			index, nested = parseInteger(s, index)
			parent.Add(*nested)
		}

		if index < len(s) && s[index] == ',' {
			index++
		}
	}

	return index, parent
}

func parseInteger(s []byte, index int) (int, *NestedInteger) {
	negtive := false
	if s[index] == '-' {
		negtive = true
		index++
	}

	integer := 0
	for ; index < len(s) && s[index] >= '0' && s[index] <= '9'; index++ {
		integer = integer*10 + int(s[index]-'0')
	}

	if negtive {
		integer = -integer
	}

	nestedInteger := &NestedInteger{}
	nestedInteger.SetInteger(integer)
	return index, nestedInteger
}
