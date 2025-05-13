func merge(intervals [][]int) (ans [][]int) {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    start, end := intervals[0][0], intervals[0][1]
    for _, e := range intervals[1:] {
        if end < e[0] {
            ans = append(ans, []int{start, end})
            start, end = e[0], e[1]
        } else if end < e[1] {
            end = e[1]
        }
    }
    ans = append(ans, []int{start, end})
    return ans
}
