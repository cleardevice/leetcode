class Solution {
    private List<List<Integer>> ans = new ArrayList<>();
    private List<Integer> t = new ArrayList<>();
    private boolean[] vis;
    private int[] nums;

    public List<List<Integer>> permuteUnique(int[] nums) {
        Arrays.sort(nums);
        this.nums = nums;
        vis = new boolean[nums.length];
        dfs(0);
        return ans;
    }

    private void dfs(int i) {
        if (i == nums.length) {
            ans.add(new ArrayList<>(t));
            return;
        }
        for (int j=0; j<nums.length; j++) {
            if (vis[j] || (j > 0 && nums[j] == nums[j-1] && !vis[j-1])) {
                continue;
            }
            vis[j] = true;
            t.add(nums[j]);
            dfs(i+1);
            t.remove(t.size()-1);
            vis[j] = false;
        }
    }
}
