func maxArea(heights []int) int {
	max := 0
	i, j := 0, len(heights) - 1
	for i < j {
		w := j-i
		if heights[i] < heights[j] {
			area := w * heights[i]
			i++
			if area > max {
				max = area
			}
		} else {
			area := w * heights[j]
			if area > max {
				max = area
			}
			j--
		}
	}
	return max
}
