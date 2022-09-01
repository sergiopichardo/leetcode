package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Right)+1, maxDepth(root.Left)+1)
}

func max(first, second int) int {
	if first > second {
		return first
	} else {
		return second
	}
}

func main() {
	// Tree 2
	// a := &TreeNode{Val: 10}
	// b := &TreeNode{Val: 5}
	// c := &TreeNode{Val: 3}
	// d := &TreeNode{Val: 1}
	// e := &TreeNode{Val: 7}
	// f := &TreeNode{Val: 6}
	// g := &TreeNode{Val: 15}
	// h := &TreeNode{Val: 13}
	// i := &TreeNode{Val: 18}

	// a.Left = b
	// a.Right = g
	// b.Left = c
	// b.Right = e
	// c.Left = d
	// g.Left = h
	// g.Right = i
	// e.Left = f

	// fmt.Println(rangeSumBST(a, 6, 10))
}
