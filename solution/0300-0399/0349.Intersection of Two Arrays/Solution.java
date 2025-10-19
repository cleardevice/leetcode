class Solution {
    public int[] intersection(int[] nums1, int[] nums2) {
        boolean[] seen = new boolean[1001];
        for (int x : nums1) {
            seen[x] = true;
        }

        int[] ans = new int[Math.min(nums1.length, nums2.length)];
        int idx = 0;

        for (int x : nums2) {
            if (seen[x]) {
                ans[idx++] = x;
                seen[x] = false;
            }
        }

        return java.util.Arrays.copyOf(ans, idx);
    }
}
