func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
    n := len(nums)
    left, right := 0, 0
    rbt := redblacktree.NewWithIntComparator()
    for right < n {
        cur := nums[right]
        right++
        if p, ok := rbt.Floor(cur); ok && cur-p.Key.(int) <= valueDiff {
            return true
        }
        if p, ok := rbt.Ceiling(cur); ok && p.Key.(int)-cur <= valueDiff {
            return true
        }
        rbt.Put(cur, struct{}{})
        if right-left > indexDiff {
            rbt.Remove(nums[left])
            left++
        }
    }
    return false
}
