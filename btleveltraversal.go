package btleveltraversal

// TreeNode basic tree node structure.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traversal(root *TreeNode, level int, result *[][]int) {
	if root == nil {
		return
	}

	if len(*result) <= level {
		arr := make([]int, 0, 10)
		*result = append(*result, arr)
	}

	arr := (*result)[level]
	(*result)[level] = append(arr, root.Val)

	traversal(root.Left, level+1, result)
	traversal(root.Right, level+1, result)
}

// LevelOrder returns a level order of the tree
func LevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0, 10)
	resPtr := &result

	traversal(root, 0, resPtr)

	return *resPtr
}
