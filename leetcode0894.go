/*
https://leetcode.com/problems/all-possible-full-binary-trees/
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
func allPossibleFBT(N int) []*TreeNode {
	numToFBT := make(map[int][]*TreeNode)
	numToFBT[0] = make([]*TreeNode, 0)
	numToFBT[1] = []*TreeNode{&TreeNode{Val: 0}}
	helper894(N, numToFBT)
	return numToFBT[N]
}

func helper894(N int, numToFBT map[int][]*TreeNode) {
	if _, exists := numToFBT[N]; !exists {
		for i := 2; i < N-1; i++ {
			helper894(i, numToFBT)
		}

		var trees []*TreeNode
		for i := 1; i < N-1; i++ {
			left := numToFBT[i]
			right := numToFBT[N-1-i]
			for _, l := range left {
				for _, r := range right {
					root := &TreeNode{Val: 0}
					root.Left = l
					root.Right = r
					trees = append(trees, root)
				}
			}
		}

		numToFBT[N] = trees
	}
}
