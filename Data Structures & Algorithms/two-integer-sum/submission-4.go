func twoSum(nums []int, target int) []int {
	storage := map[int]int{}
	for j, n := range nums {
		diff := target - n
		if i, ok := storage[diff]; ok {
			return []int{i,j}
		}
		storage[n] = j
	}
	return nil
}