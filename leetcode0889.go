/*
LeetCode 889: https://leetcode.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal/
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
func constructFromPrePost(pre []int, post []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}

	rootValue := pre[0]
	leftCount := partitionTree(pre[1:], post[0:len(post)-1])
	return &TreeNode{
		Val:   rootValue,
		Left:  constructFromPrePost(pre[1:1+leftCount], post[:leftCount]),
		Right: constructFromPrePost(pre[1+leftCount:], post[leftCount:len(post)-1]),
	}
}

func partitionTree(pre []int, post []int) int {
	maxJ := 0
	for i, val1 := range pre {
		for j, val2 := range post {
			if val1 == val2 {
				if maxJ < j {
					maxJ = j
				}

				break
			}
		}

		if i == maxJ {
			return maxJ + 1
		}
	}

	return 0
}
