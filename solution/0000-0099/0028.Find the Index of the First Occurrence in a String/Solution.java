class Solution {
    public int strStr(String haystack, String needle) {
        int lenP = haystack.length();
        int lenQ = needle.length();
        if (lenQ == 0) {
            return 0;
        }
        if (lenQ > lenP) {
            return -1;
        }
        int p = 0;
        int q = 0;
        while (p < lenP - lenQ + 1) {
            if (haystack.charAt(p + q) == needle.charAt(q)) {
                if (++q == lenQ) {
                    return p;
                }
                continue;
            }
            p++;
            q = 0;
        }
        return -1;
    }
}
