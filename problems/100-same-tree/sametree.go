package sametree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return preorderTraversalRecursive(true, p, q)
}

func preorderTraversalRecursive(isSame bool, p *TreeNode, q *TreeNode) bool {
	if p != nil && q != nil {
		isSame = isSame && p.Val == q.Val
		isSame = isSame && preorderTraversalRecursive(isSame, p.Left, q.Left)
		isSame = isSame && preorderTraversalRecursive(isSame, p.Right, q.Right)
	}
	if p == nil && q != nil {
		return false
	}
	if p != nil && q == nil {
		return false
	}
	return isSame
}
