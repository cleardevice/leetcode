func removeDuplicates(nums []int) int {
    k := 0
    for _, v := range nums {
        if k == 0 || v != nums[k-1] {
            nums[k] = v
            k++
        }
    }
    return k
}
