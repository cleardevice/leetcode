class Solution {
    List<String> result = new ArrayList<>();

    public List<String> generateParenthesis(int n) {
        process("", 0, 0, n);
        return result;
    }

    public void process(String string, int opened, int closed, int max) {
        if (string.length() == max * 2) {
            result.add(string);
        }
        if (opened < max) {
            process(string + "(", opened+1, closed, max);
        }
        if (closed < opened) {
            process(string + ")", opened, closed+1, max);
        }
    }
}
