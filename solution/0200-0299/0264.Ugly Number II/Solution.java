class Solution {
    public int nthUglyNumber(int n) {
        PriorityQueue<Long> q = new PriorityQueue<>();
        Set<Long> vis = new HashSet<>();
        q.offer(1L);
        vis.add(1L);
        long ans = 0;
        int[] f = new int[] {2, 3, 5};
        while (n-- > 0) {
            ans = q.poll();
            for (int v : f) {
                long next = v * ans;
                if (vis.add(next)) {
                    q.offer(next);
                }
            }
        }
        return (int) ans;
    }
}
