function merge(intervals: number[][]): number[][] {
    intervals.sort((a, b) => a[0] - b[0]);
    const ans: number[][] = [];
    let [start, end] = intervals[0];
    for (const [s, e] of intervals.slice(1)) {
        if (end < s) {
            ans.push([start, end]);
            [start, end] = [s, e];
        } else {
            end = Math.max(end, e);
        }
    }
    ans.push([start, end]);
    return ans;
};
