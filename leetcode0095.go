/*
LeetCode 95: https://leetcode.com/problems/unique-binary-search-trees-ii/
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
func generateTrees(n int) []*TreeNode {
	return helper95(1, n)
}

func helper95(low, upper int) []*TreeNode {
	if low > upper {
		return make([]*TreeNode, 0)
	} else if low == upper {
		result := make([]*TreeNode, 1)
		result[0] = &TreeNode{
			Val: low,
		}

		return result
	} else {
		result := make([]*TreeNode, 0)
		for i := low; i <= upper; i++ {
			left := helper95(low, i-1)
			right := helper95(i+1, upper)
			result = addTrees95(result, left, right, i)
		}

		return result
	}
}

func addTrees95(trees, left, right []*TreeNode, rootVal int) []*TreeNode {
	if len(left) == 0 {
		for _, r := range right {
			root := &TreeNode{
				Val:   rootVal,
				Right: r,
			}

			trees = append(trees, root)
		}
	} else if len(right) == 0 {
		for _, l := range left {
			root := &TreeNode{
				Val:  rootVal,
				Left: l,
			}

			trees = append(trees, root)
		}
	} else {
		for _, l := range left {
			for _, r := range right {
				root := &TreeNode{
					Val:   rootVal,
					Left:  l,
					Right: r,
				}

				trees = append(trees, root)
			}
		}
	}

	return trees
}
