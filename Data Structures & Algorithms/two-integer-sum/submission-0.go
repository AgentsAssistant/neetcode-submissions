func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)
	for i, n := range nums {
		diff := target - n
		if j, ok := seen[diff]; ok {
			return _sort(i, j)
		}
		seen[n] = i
	}
	return nil
}

func _sort(i, j int) ([]int) {
	if i < j {
		return []int{i, j}
	}
	return []int{j, i}
}