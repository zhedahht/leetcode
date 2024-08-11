/*
LeetCode 437: https://leetcode.com/problems/path-sum-iii/description/
*/

package leetcode

// The function name should be pathSum. It's renamed to avoid conflicts.
func pathSum437(root *TreeNode, targetSum int) int {
	sums := []int{0}
	return helper437(root, targetSum, &sums)
}

func helper437(root *TreeNode, targetSum int, sums *[]int) int {
	if root == nil {
		return 0
	}

	count := 0
	sum := (*sums)[len(*sums)-1] + root.Val
	for _, s := range *sums {
		if s+targetSum == sum {
			count++
		}
	}

	*sums = append(*sums, sum)

	leftCount := helper437(root.Left, targetSum, sums)
	rightCount := helper437(root.Right, targetSum, sums)

	*sums = (*sums)[:len(*sums)-1]
	return count + leftCount + rightCount
}
