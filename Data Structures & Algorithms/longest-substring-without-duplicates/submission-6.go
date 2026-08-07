func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	seen := make(map[byte]int)
	i, j := 0, 0
	res := 0
	for ; j < len(s); j++ {
		if k, ok := seen[s[j]]; ok && k >= i {
			i = k+1
		}
		seen[s[j]] = j
		if j - i > res {
			res = j - i
		}
	}
	return res+1
}
