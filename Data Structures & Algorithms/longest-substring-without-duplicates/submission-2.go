func lengthOfLongestSubstring(s string) int {
	l := len(s)
	if l < 2 {
		return l
	}
	seen := map[byte]int{}
	res := 1
	i, j := 0, 0
	for j < l {
		if k, ok := seen[s[j]]; ok && k >= i {
			if j-i > res {
				res = j-i
			}
			i = k+1
		}
		seen[s[j]] = j
		j++
	}
	if j-i > res {
		res = j-i
	}
	return res
}
