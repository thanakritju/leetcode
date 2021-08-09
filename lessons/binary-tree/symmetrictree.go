package binarytree

func isSymmetric(root *TreeNode) bool {
	return topDown(true, root.Left, root.Right)
}

func topDown(answer bool, p *TreeNode, q *TreeNode) bool {
	if p != nil && q != nil {
		answer = answer && p.Val == q.Val
		answer = answer && topDown(answer, p.Left, q.Right)
		answer = answer && topDown(answer, p.Right, q.Left)
	}
	if p == nil && q != nil {
		return false
	}
	if p != nil && q == nil {
		return false
	}
	return answer
}
