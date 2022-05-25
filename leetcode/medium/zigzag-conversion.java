class Solution {
    public String convert(String s, int numRows) {
        if (numRows == 1) {
            return s;
        }
        
        List<Character>[] answer = new List[numRows];
        for (int i = 0; i < numRows; i++) {
            answer[i] = new ArrayList<>();
        }
        
        int cur = 0;
        int row = 0;
        boolean down = true;
        while (cur != s.length()) {
            answer[row].add(s.charAt(cur));
            cur++;
            if (down) {
                row++;
            } else {
                row--;
            }
            
            if (row == numRows) {
                down = false;
                row = numRows-2;
            } else if (row == -1) {
                down = true;
                row = 1;
            }
            
        }
        
        StringBuilder sb = new StringBuilder();
        for (List<Character> list : answer) {
            for (Character c : list) {
                sb.append(c);
            }
        }
        return sb.toString();
    }
}
