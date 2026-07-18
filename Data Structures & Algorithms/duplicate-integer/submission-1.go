func hasDuplicate(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	seen := make(map[int]struct{})
	for _, n := range nums {
		if _, ok := seen[n]; ok {
			return true
		}
		seen[n] = struct{}{}
	}
	return false
}
