class Solution {
    public String convert(String s, int numRows) {
        if (numRows == 1) {
            return s;
        }
        StringBuilder[] lines = new StringBuilder[numRows];
        Arrays.setAll(lines, l -> new StringBuilder());
        int i = 0, step = 1;
        for (char c : s.toCharArray()) {
            lines[i].append(c);
            i += step;
            if (i == 0 || i == numRows-1) {
                step = -step;
            }
        }
        return String.join("", lines);
    }
}
