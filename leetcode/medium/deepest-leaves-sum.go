/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
    deepest, answer := 0, 0
    dfs(root, 0, &deepest, &answer)
    return answer
}

func dfs(parent *TreeNode, depth int, deepest, answer *int) {   
    if parent.Left == nil && parent.Right == nil {
        if depth == *deepest {
            *answer += parent.Val
        } 
        
        if depth > *deepest {
            *deepest, *answer = depth, 0
            *answer = parent.Val
        }
    }

    if parent.Left != nil {
        dfs(parent.Left, depth + 1, deepest, answer)
    }

    if parent.Right != nil {
        dfs(parent.Right, depth + 1, deepest, answer)
    }
}
