class Solution {
    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> map = new HashMap<>();
        for (int i=0; i<nums.length; i++) {
            int curr = nums[i];
            if (map.containsKey(target-curr)) {
                return new int[]{map.get(target-curr), i};
            }
            map.put(curr, i);
        }
        throw new RuntimeException();
    }
}
