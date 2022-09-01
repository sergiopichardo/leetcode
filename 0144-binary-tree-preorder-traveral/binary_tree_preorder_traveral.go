package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := make([]int, 0)
	return helper(root, &result)
}

func helper(node *TreeNode, values *[]int) []int {
	if node == nil {
		return *values
	}

	*values = append(*values, node.Val)

	helper(node.Left, values)
	helper(node.Right, values)

	return *values
}
