/*
LeetCode 589: https://leetcode.com/problems/n-ary-tree-preorder-traversal/description/
*/

package leetcode

// The name should be Node. It's renamed to avoid conflicts.
type Node589 struct {
	Val      int
	Children []*Node589
}

func preorder(root *Node589) []int {
	result := make([]int, 0)
	preorder_helper(root, &result)
	return result
}

func preorder_helper(node *Node589, result *[]int) {
	if node == nil {
		return
	}

	*result = append(*result, node.Val)
	for _, child := range node.Children {
		preorder_helper(child, result)
	}
}
