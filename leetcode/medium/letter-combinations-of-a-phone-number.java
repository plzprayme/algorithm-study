class Solution {
    
    private static Map<Character, String> map = Map.of(
        '2', "abc",
        '3', "def",
        '4', "ghi",
        '5', "jkl",
        '6', "mno",
        '7', "pqrs",
        '8', "tuv",
        '9', "wxyz"
    );
    
    public List<String> letterCombinations(String digits) {
        List<String> result = new ArrayList<>();
        if (!digits.equals("")) {
            combination(0, digits.length(), "", digits, result);
        }
        return result;
    }
    
    public void combination(int curSize, int maxSize, String curDigits, String digits, List<String> list) {
        if (curSize == maxSize) {
            list.add(curDigits);
            return;
        } 
        
        String s = map.get(digits.charAt(curSize));
        for (char c : s.toCharArray()) {
            combination(curSize + 1, maxSize, curDigits + c, digits, list);
        }
    }
}
