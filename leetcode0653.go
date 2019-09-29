/*
LeetCode 653: https://leetcode.com/problems/two-sum-iv-input-is-a-bst/
*/

package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
	node1, node2 := root, root
	stack1, stack2 := make([]*TreeNode, 0), make([]*TreeNode, 0)

	getNext := func() *TreeNode {
		for node1 != nil || len(stack1) > 0 {
			for node1 != nil {
				stack1 = append(stack1, node1)
				node1 = node1.Left
			}

			result := stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
			node1 = result.Right
			return result
		}

		return nil
	}

	getPrev := func() *TreeNode {
		for node2 != nil || len(stack2) > 0 {
			for node2 != nil {
				stack2 = append(stack2, node2)
				node2 = node2.Right
			}

			result := stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
			node2 = result.Left
			return result
		}

		return nil
	}

	first, last := getNext(), getPrev()
	for first != last && first != nil && last != nil {
		if first.Val+last.Val == k {
			return true
		}

		if first.Val+last.Val > k {
			last = getPrev()
		} else {
			first = getNext()
		}
	}

	return false
}
