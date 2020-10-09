/*
LeetCode 230: https://leetcode.com/problems/kth-smallest-element-in-a-bst/
*/

package leetcode

func kthSmallest(root *TreeNode, k int) int {
	stack := make([]*TreeNode, 0)
	cur := root
	i := 0
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		i++
		if i == k {
			return cur.Val
		}

		cur = cur.Right
	}

	return -1
}
