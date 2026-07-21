func hasDuplicate(nums []int) bool {
	seen := make(map[int]struct{})
	for _, num := range nums {
		if _, isFound := seen[num]; isFound {
			return true
		}
		seen[num] = struct{}{}
	}
	return false
}
