func maxProfit(prices []int) int {
	min := 100
	res := 0
	for _, p := range prices {
		if profit := p - min; profit > res {
			res = profit
		}
		if p < min {
			min = p
		}
	}
	return res
}
