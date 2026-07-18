func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := make(map[rune]int)
	for _, c := range s {
		sMap[c] += 1
	}
	for _, c := range t {
		if n, ok := sMap[c]; !ok || n < 1 {
			return false
		}
		sMap[c] -= 1
	}
	return true
}
