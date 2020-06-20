/*
LeetCode 993: https://leetcode.com/problems/cousins-in-binary-tree/
*/

package leetcode

func isCousins(root *TreeNode, x int, y int) bool {
	parentX, heightX := helper993(root, 0, x)
	parentY, heightY := helper993(root, 0, y)
	return parentX != parentY && heightX == heightY
}

func helper993(root *TreeNode, h, x int) (*TreeNode, int) {
	if root == nil || root.Val == x {
		return nil, 0
	}

	if root.Left != nil && root.Left.Val == x {
		return root, h + 1
	}

	if root.Right != nil && root.Right.Val == x {
		return root, h + 1
	}

	leftH, leftX := helper993(root.Left, h+1, x)
	if leftH != nil {
		return leftH, leftX
	}

	return helper993(root.Right, h+1, x)
}
