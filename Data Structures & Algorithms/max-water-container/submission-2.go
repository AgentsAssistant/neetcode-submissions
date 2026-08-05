func maxArea(heights []int) int {
	res := 0
	i, j := 0, len(heights) - 1
	for i < j {
		area := min(heights[i], heights[j]) * (j-i)
		if area > res {
			res = area
		}
		if heights[i] < heights[j] {
			i++
		} else {
			j--
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}