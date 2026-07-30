func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	n := len(nums)-1
	m := n-1
	for i := 0; i < m; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i+1
		k := n
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum < 0 {
				j++
				continue
			} 
			if sum > 0 {
				k--
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
