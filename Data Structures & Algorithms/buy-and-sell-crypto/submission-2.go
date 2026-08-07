func maxProfit(prices []int) int {
	res := 0
	min := 1000
	for _, p := range prices {
		if p - min > res {
			res = p - min
		}
		if p < min {
			min = p
		}
	}
	return res
}
