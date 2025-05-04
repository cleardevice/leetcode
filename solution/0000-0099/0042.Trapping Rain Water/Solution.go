func trap(height []int) int {
    n := len(height)
    left := make([]int, n)
    right := make([]int, n)
    left[0], right[n-1] = height[0], height[n-1]
    for i := 1; i < n; i++ {
        left[i] = max(left[i-1], height[i])
        right[n-i-1] = max(right[n-i], height[n-i-1])
    }
    ans := 0
    for i := 0; i < n; i++ {
        ans += min(left[i], right[i])-height[i]
    }
    return ans
}
