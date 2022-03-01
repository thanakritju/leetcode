package pathsumiii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	s, v := pathSumRec(root, 0, []int{}, []*TreeNode{})
	count := 0
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			pathSum := s[j] - s[i] + v[i].Val
			if pathSum == targetSum && findPath(v[i], v[j]) {
				count += 1
			}
		}
	}
	return count
}

func findPath(node *TreeNode, target *TreeNode) bool {
	if node == nil {
		return false
	}
	if node == target {
		return true
	}
	if findPath(node.Left, target) {
		return true
	}
	if findPath(node.Right, target) {
		return true
	}
	return false
}

func pathSumRec(node *TreeNode, sum int, s []int, v []*TreeNode) ([]int, []*TreeNode) {
	if node != nil {
		sum = sum + node.Val
		s = append(s, sum)
		v = append(v, node)
		s, v = pathSumRec(node.Left, sum, s, v)
		s, v = pathSumRec(node.Right, sum, s, v)
	}
	return s, v
}
