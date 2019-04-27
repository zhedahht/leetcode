/*
LeetCode 23: https://leetcode.com/problems/merge-k-sorted-lists/
*/

package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	return helper23(lists, 0, len(lists)-1)
}

func helper23(lists []*ListNode, lower, upper int) *ListNode {
	if lower > upper {
		return nil
	} else if lower == upper {
		return lists[lower]
	}

	mid := (lower + upper) / 2
	list1 := helper23(lists, lower, mid)
	list2 := helper23(lists, mid+1, upper)

	// NOTE: mergeTwoLists is in leetcode0021.go
	return mergeTwoLists(list1, list2)
}
