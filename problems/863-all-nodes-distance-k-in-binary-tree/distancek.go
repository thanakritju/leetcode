package distancek

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	target := searchBST(root, target.Val)
	return []int{}
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root != nil {
		if root.Val == val {
			return root
		}
		if temp := searchBST(root.Left, val); temp != nil {
			return temp
		}
		if temp := searchBST(root.Right, val); temp != nil {
			return temp
		}
	}
	return nil
}
