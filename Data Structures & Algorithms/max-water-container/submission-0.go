func maxArea(heights []int) int {
	max := 0
	i, j := 0, len(heights)-1
	for i < j {
		height := heights[i]
		width := j-i
		if heights[j] < height {
			height = heights[j]
			j--
		} else {
			i++
		}
		area := width*height
		if area > max {
			max = area
		}
	}
	return max
}
