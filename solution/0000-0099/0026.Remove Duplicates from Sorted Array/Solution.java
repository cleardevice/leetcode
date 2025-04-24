class Solution {
    public int removeDuplicates(int[] nums) {
        int k = 0;
        for (int v : nums) {
            if (k == 0 || v != nums[k-1]) {
                nums[k++] = v;
            }
        }
        return k;
    }
}
