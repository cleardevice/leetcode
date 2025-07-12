func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
    rbt := redblacktree.NewWithIntComparator()
    for i, cur := range nums {
        if p, ok := rbt.Floor(cur); ok && cur-p.Key.(int) <= valueDiff {
            return true
        }
        if p, ok := rbt.Ceiling(cur); ok && p.Key.(int)-cur <= valueDiff {
            return true
        }
        rbt.Put(cur, struct{}{})
        if i >= indexDiff {
            rbt.Remove(nums[i-indexDiff])
        }
    }
    return false
}
