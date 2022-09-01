/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rangeSumBST(root *TreeNode, low int, high int) int {
	total := 0
	return sum(root, low, high, total)
}

func sum(node *TreeNode, low int, high int, total int) int {
	if node == nil {
		return total
	}

	if node.Val >= low && node.Val <= high {
		total += node.Val
	}

	total = sum(node.Left, low, high, total)
	total = sum(node.Right, low, high, total)

	return total
}