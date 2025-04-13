function lengthOfLongestSubstring(s: string): number {
    const cnt = new Map<string, number>();
    let ans = 0;
    for (let l=0, r=0; r < s.length; ++r) {
        cnt.set(s[r], (cnt.get(s[r]) || 0) + 1);
        while (cnt.get(s[r])! > 1) {
            cnt.set(s[l], cnt.get(s[l])! - 1);
            ++l;
        }
        ans = Math.max(ans, r-l+1);
    }
    return ans;
};
