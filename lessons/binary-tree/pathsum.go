package binarytree

func hasPathSum(root *TreeNode, targetSum int) bool {
	ans := false
	subSum := targetSum - root.Val

	if subSum == 0 && root.Left == nil && root.Right == nil {
		return true
	}

	if root.Left != nil {
		ans = ans || hasPathSum(root.Left, subSum)
	}
	if root.Right != nil {
		ans = ans || hasPathSum(root.Right, subSum)
	}
	return ans
}
