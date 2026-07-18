func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	arr := make([]int, 26)
	for _, c := range s {
		arr[c-'a'] += 1
	}
	for _, c := range t {
		arr[c-'a'] -= 1
	}
	for _, c := range arr {
		if c != 0 {
			return false
		}
	}
	return true
}
