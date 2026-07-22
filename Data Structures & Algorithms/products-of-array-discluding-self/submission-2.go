func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = 1
	for i := 0; i < len(nums)-1; i++ {
		res[i+1] = res[i] * nums[i]
	}
	postProd := 1
	for j := len(nums) - 1; j >= 0; j-- {
		res[j] = postProd * res[j]
		postProd *= nums[j]
	}
	return res
}
