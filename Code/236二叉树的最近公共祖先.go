package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isIn(root, target *TreeNode) bool {
	if root == nil {
		return false
	}
	if root == target {
		return true
	}
	return isIn(root.Left, target) || isIn(root.Right, target)

}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || (isIn(root.Left, p) && isIn(root.Right, q)) || (isIn(root.Left, q) && isIn(root.Right, p)) {
		return root
	}
	if isIn(p, q) {
		return p
	} else if isIn(q, p) {
		return q
	}
	if isIn(root.Left, p) {
		return lowestCommonAncestor(root.Left, p, q)
	}
	return lowestCommonAncestor(root.Right, p, q)
}
