/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */

class Solution {
    
    TreeNode answer;
    TreeNode t;
    
    public final TreeNode getTargetCopy(final TreeNode original, final TreeNode cloned, final TreeNode target) {
        t = target;
        dfs(original, cloned);
        return answer;
    }
    
    private void dfs(TreeNode origin, TreeNode cloned) {
        if (Objects.isNull(origin) || Objects.isNull(cloned)) {
            return;
        }
        
        if (origin.equals(t)) {
            answer = cloned;
            return;
        }
        
        
        dfs(origin.left, cloned.left);
        dfs(origin.right, cloned.right);
    }
}
