package distancek

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	ans, _ := findKdistance(root, target, k, []int{})
	return ans
}

func findKdistance(root *TreeNode, target *TreeNode, k int, ans []int) ([]int, int) {
	if root == nil {
		return ans, -1
	}

	if root.Val == target.Val {
		ans = findKdistanceNodesDown(root, k, ans)
		return ans, 0
	}

	ans, dl := findKdistance(root.Left, target, k, ans)
	if dl != -1 {
		if dl+1 == k {
			ans = append(ans, root.Val)
		} else {
			ans = findKdistanceNodesDown(root.Right, k-dl-2, ans)
		}
		return ans, dl + 1
	}

	ans, dr := findKdistance(root.Right, target, k, ans)
	if dr != -1 {
		if dr+1 == k {
			ans = append(ans, root.Val)
		} else {
			ans = findKdistanceNodesDown(root.Left, k-dr-2, ans)
		}
		return ans, dr + 1
	}

	return ans, -1
}

func findKdistanceNodesDown(node *TreeNode, k int, ans []int) []int {
	if node == nil {
		return ans
	}
	if k == 0 {
		ans = append(ans, node.Val)
	}
	ans = findKdistanceNodesDown(node.Left, k-1, ans)
	ans = findKdistanceNodesDown(node.Right, k-1, ans)
	return ans
}
