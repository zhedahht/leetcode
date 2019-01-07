/*
LeetCode 971: https://leetcode.com/problems/flip-binary-tree-to-match-preorder-traversal/
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
func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	result := make([]int, 0)
	result, ok := flipCore(root, voyage, 0, len(voyage)-1, result)
	if ok {
		return result
	}

	return []int{-1}
}

func flipCore(root *TreeNode, voyage []int, start int, end int, result []int) ([]int, bool) {
	if root == nil {
		return result, start > end
	} else {
		if root.Val != voyage[start] {
			return result, false
		}

		if root.Left == nil && root.Right == nil {
			return result, true
		}

		if root.Left == nil {
			return flipCore(root.Right, voyage, start+1, end, result)
		}

		if root.Right == nil {
			return flipCore(root.Left, voyage, start+1, end, result)
		}

		if end-start < 2 {
			return result, false
		}

		if root.Left.Val == voyage[start+1] {
			i := start + 2
			for ; i <= end; i++ {
				if root.Right.Val == voyage[i] {
					break
				}
			}

			if i > end {
				return result, false
			}

			result, ok1 := flipCore(root.Left, voyage, start+1, i-1, result)
			result, ok2 := flipCore(root.Right, voyage, i, end, result)
			return result, ok1 && ok2
		}

		if root.Right.Val == voyage[start+1] {
			i := start + 2
			for ; i <= end; i++ {
				if root.Left.Val == voyage[i] {
					break
				}
			}

			if i > end {
				return result, false
			}

			result = append(result, root.Val)
			result, ok1 := flipCore(root.Right, voyage, start+1, i-1, result)
			result, ok2 := flipCore(root.Left, voyage, i, end, result)
			return result, ok1 && ok2
		}

		return result, false
	}
}
