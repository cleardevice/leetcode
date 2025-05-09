class Solution {
    public int lengthOfLongestSubstring(String s) {
        int[] cnt = new int[128];
        int ans = 0;
        for (int l=0, r=0; r < s.length(); ++r) {
            char c = s.charAt(r);
            ++cnt[c];
            while (cnt[c] > 1) {
                --cnt[s.charAt(l++)];
            }
            ans = Math.max(ans, r-l+1);
        }
        return ans;
    }
}
