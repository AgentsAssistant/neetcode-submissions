func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers) - 1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			break
		}
		if sum > target {
			j--
			continue
		}
		if sum < target {
			i++
			continue
		}
	}
	return []int{i+1, j+1}
}
