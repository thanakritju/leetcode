package pathsumiii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	return findPath(root, &[]*TreeNode{}, targetSum)
}

func findPath(node *TreeNode, path *[]*TreeNode, target int) int {
	if node == nil {
		return 0
	}
	ans := 0
	*path = append(*path, node)
	ans += findPath(node.Left, path, target)
	ans += findPath(node.Right, path, target)

	sum := 0
	for i := len(*path) - 1; i > -1; i-- {
		sum += (*path)[i].Val
		if sum == target {
			ans += 1
		}
	}

	*path = (*path)[:len(*path)-1]

	return ans
}
