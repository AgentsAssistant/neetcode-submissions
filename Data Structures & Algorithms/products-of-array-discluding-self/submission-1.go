func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	res[0] = 1
	for i := 0; i < n-1; i++ {
		res[i+1] = res[i] * nums[i]
	}
	postProd := 1
	for i := n-1; i > -1; i-- {
		res[i] *= postProd
		postProd *= nums[i]
	}
	return res
}
