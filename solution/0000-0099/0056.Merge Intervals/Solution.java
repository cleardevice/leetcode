class Solution {
    public int[][] merge(int[][] intervals) {
        Arrays.sort(intervals, Comparator.comparingInt(a -> a[0]));
        int start = intervals[0][0], end = intervals[0][1];
        List<int[]> ans = new ArrayList<>();
        for (int i = 1; i < intervals.length; i++) {
            int s = intervals[i][0], e = intervals[i][1];
            if (end < s) {
                ans.add(new int[] { start, end });
                start = s;
                end = e;
            } else {
                end = Math.max(end, e);
            }
        }
        ans.add(new int[] { start, end });
        return ans.toArray(new int[ans.size()][]);
    }
}
