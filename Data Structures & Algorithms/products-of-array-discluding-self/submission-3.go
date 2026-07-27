func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = 1
	for i := 0; i < len(nums)-1; i++ {
		res[i+1] = res[i]*nums[i]
	}
	postProd := 1
	for i := len(nums)-1; i >= 0; i-- {
		res[i] *= postProd
		postProd *= nums[i]
	}
	return res
}
