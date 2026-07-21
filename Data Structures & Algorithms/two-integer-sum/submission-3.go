func twoSum(nums []int, target int) []int {
	first := map[int]int{}
	for j, num := range nums {
		diff := target - num
		if i, found := first[diff]; found {
			return []int{i, j}
		}
		first[num] = j
	}
	return nil
}