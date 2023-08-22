package main

func pathSum(root *TreeNode, targetSum int) int {
	res := 0
	var dfs func(cur *TreeNode, target, status int)
	dfs = func(cur *TreeNode, target, status int) {
		if cur == nil {
			return
		}
		if cur.Val == target {
			res++
		}
		dfs(cur.Left, target-cur.Val, 1)
		dfs(cur.Right, target-cur.Val, 1)

		if status == 0 {
			dfs(cur.Left, target, 0)
			dfs(cur.Right, target, 0)
		}

	}
	dfs(root, targetSum, 0)
	return res
}
