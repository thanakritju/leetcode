package pathsumii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	return getAllArrays(root, [][]int{}, []int{}, targetSum)
}

func getAllArrays(root *TreeNode, arrays [][]int, array []int, targetSum int) [][]int {
	if root != nil {
		array = append(array, root.Val)
		if root.Left == nil && root.Right == nil && sumArray(array) == targetSum {
			copiedArray := make([]int, len(array))
			copy(copiedArray, array)
			arrays = append(arrays, copiedArray)
		} else {
			arrays = getAllArrays(root.Left, arrays, array, targetSum)
			arrays = getAllArrays(root.Right, arrays, array, targetSum)
		}
	}
	return arrays
}

func sumArray(array []int) int {
	sum := 0
	for _, each := range array {
		sum += each
	}
	return sum
}
