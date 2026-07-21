func productExceptSelf(nums []int) []int {
	length := len(nums)
	pre := make([]int, length)
	pre[0] = 1
	suf := make([]int, length)
	suf[length-1] = 1
	for i := 0; i < length-1; i++ {
		pre[i+1] = nums[i]*pre[i]
	}
	for i := length-1; i > 0; i-- {
		suf[i-1] = nums[i]*suf[i]
	}
	for i, _ := range pre {
		pre[i] *= suf[i]
	}
	return pre
}
