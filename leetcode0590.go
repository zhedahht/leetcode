/*
LeetCode 590: https://leetcode.com/problems/n-ary-tree-postorder-traversal/description/
*/

package leetcode

func postorder(root *Node589) []int {
	result := make([]int, 0)
	postorder_helper(root, &result)
	return result
}

func postorder_helper(node *Node589, result *[]int) {
	if node == nil {
		return
	}

	for _, child := range node.Children {
		postorder_helper(child, result)
	}

	*result = append(*result, node.Val)
}
