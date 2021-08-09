package binarytree

func maxDepth(root *TreeNode) int {
	return maxDepthTopDown(root, 0, 1)
}

func maxDepthTopDown(root *TreeNode, answer int, depth int) int {
	if root != nil {
		if depth > answer {
			answer = depth
		}
		answer = maxDepthTopDown(root.Left, answer, depth+1)
		answer = maxDepthTopDown(root.Right, answer, depth+1)
	}
	return answer
}
