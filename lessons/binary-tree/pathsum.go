package binarytree

func hasPathSum(root *TreeNode, targetSum int) bool {
	return hasPathSumTopDown(root, false, 0, targetSum)
}

func hasPathSumTopDown(root *TreeNode, answer bool, currentSum int, targetSum int) bool {
	if root != nil {
		currentSum += root.Val
		answer = answer || (currentSum == targetSum && root.Left == nil && root.Right == nil)
		answer = answer || hasPathSumTopDown(root.Left, answer, currentSum, targetSum)
		answer = answer || hasPathSumTopDown(root.Right, answer, currentSum, targetSum)
	}
	return answer
}
