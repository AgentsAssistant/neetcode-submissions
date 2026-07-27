func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums) - 2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum > 0 {
				k--
				continue
			}
			if sum < 0 {
				j++
				continue
			}
			res = append(res, []int{nums[i], nums[j], nums[k]})
			j++
			k--
			for j < k && nums[j] == nums[j-1] {
				j++
			}
		}
	}
	return res
}
