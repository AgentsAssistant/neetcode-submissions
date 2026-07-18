func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	arr := [26]int{}
	for i, c := range s {
		arr[c-'a']++
		arr[t[i]-'a']--
	}
	for _, c := range arr {
		if c != 0 {
			return false
		}
	}
	return true
}
