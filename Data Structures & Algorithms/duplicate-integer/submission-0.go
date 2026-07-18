func hasDuplicate(nums []int) bool {
    n := len(nums)
    if n < 2 {
        return false
    }
    set := map[int]bool{}
    for i := 0; i < n; i++ {
        if set[nums[i]] {
            return true
        }
        set[nums[i]] = true
    }
    return false
}
