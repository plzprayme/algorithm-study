class Solution {
    
    private int maxSize;
    private int maxSum;
    private final boolean[] used = new boolean[10];
    private List<List<Integer>> answer;
    private int[] selected;
    public List<List<Integer>> combinationSum3(int k, int n) {
        maxSize = k;
        maxSum = n;
        answer = new ArrayList<>();
        selected = new int[k];
        combination(0, 0, 1, new LinkedList<>());
        return answer;
    }
    
    private void combination(int curSize, int curSum, int start, LinkedList<Integer> curList) {
        
        if (curSize == maxSize && curSum == maxSum) {
            
            List<Integer> list = new ArrayList<>();
            for (int i = 0; i < maxSize; i++) {
                list.add(selected[i]);
            }
            answer.add(list);
            
            return;    
        } else if (curSize == maxSize || curSum > maxSum) {
            return;
        }
        
        
        for (int i = start; i < 10; i++) {
            if (used[i]) continue;
            used[i] = true;
            selected[curSize] = i; 
            combination(curSize + 1, curSum + i, i + 1, curList);
            selected[curSize] = -1;
            used[i] = false;
        }
        
    }
}
