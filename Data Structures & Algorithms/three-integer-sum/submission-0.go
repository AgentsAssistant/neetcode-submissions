func threeSum(nums []int) [][]int {
	index := map[int]int{}
	for i, n := range nums {
		index[n] = i
	}
	res := [][]int{}
	output := map[[3]int]bool{}
	for i, n1 := range nums {
		for j, n2 := range nums {
			if i == j {
				continue
			}
			diff := 0 - n1 - n2
			if k, ok := index[diff]; ok && i != k && j != k {
				trip := sort3(n1, n2, diff)
				if _, ok := output[trip]; !ok {
					output[trip] = true
					res = append(res, []int{n1, n2, diff})
				}
			}
		}
	}
	return res
}

func sort3(a, b, c int) ([3]int) {
	if a > b {
		a, b = b, a
	}
	if b > c {
		b, c = c, b
		if a > b {
			a, b = b, a
		}
	}
	return [3]int{a, b, c}
}
