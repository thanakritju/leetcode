package binarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	arr := preorderTraversalRecursive([]int{}, root)
	return arr
}

func preorderTraversalRecursive(arr []int, root *TreeNode) []int {
	if root != nil {
		arr = append(arr, root.Val)
		arr = preorderTraversalRecursive(arr, root.Left)
		arr = preorderTraversalRecursive(arr, root.Right)
	}
	return arr
}

func inorderTraversal(root *TreeNode) []int {
	arr := inorderTraversalRecursive([]int{}, root)
	return arr
}

func inorderTraversalRecursive(arr []int, root *TreeNode) []int {
	if root != nil {
		arr = inorderTraversalRecursive(arr, root.Left)
		arr = append(arr, root.Val)
		arr = inorderTraversalRecursive(arr, root.Right)
	}
	return arr
}

func postorderTraversal(root *TreeNode) []int {
	arr := postorderTraversalRecursive([]int{}, root)
	return arr
}

func postorderTraversalRecursive(arr []int, root *TreeNode) []int {
	if root != nil {
		arr = postorderTraversalRecursive(arr, root.Left)
		arr = postorderTraversalRecursive(arr, root.Right)
		arr = append(arr, root.Val)
	}
	return arr
}

func levelOrder(root *TreeNode) [][]int {
	arr := levelOrderRecursive([][]int{}, root, 0)
	return arr
}

func levelOrderRecursive(arr [][]int, root *TreeNode, depth int) [][]int {
	if root != nil {
		if len(arr) <= depth {
			arr = append(arr, []int{})
		}
		arr[depth] = append(arr[depth], root.Val)
		arr = levelOrderRecursive(arr, root.Left, depth+1)
		arr = levelOrderRecursive(arr, root.Right, depth+1)
	}
	return arr
}
