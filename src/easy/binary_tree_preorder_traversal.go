package easy

import . "github.com/johnazedo/leetcode/src"

/*
Difficult: Easy
Link: https://leetcode.com/problems/binary-tree-preorder-traversal/
*/
func preorderTraversal(root *TreeNode) []int {
	// TODO: Do this recursively
	stack := []*TreeNode{root}
	var result []int
	// N is the number of nodes and H is the height of the tree
	// Time: O(N)
	// Space: O(H)
	for len(stack) > 0 {
		n := len(stack) - 1
		root = stack[n]
		stack = stack[:n]
		if root != nil {
			result = append(result, root.Val)
			if root.Right != nil {
				stack = append(stack, root.Right)
			}
			if root.Left != nil {
				stack = append(stack, root.Left)
			}
		}
	}

	return result
}
