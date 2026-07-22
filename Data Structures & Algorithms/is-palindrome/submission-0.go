func isPalindrome(s string) bool {
	var alphaNumeric [128]bool
	for c := '0'; c <= '9'; c++ { alphaNumeric[c] = true }
	for c := 'A'; c <= 'Z'; c++ { alphaNumeric[c] = true }
	for c := 'a'; c <= 'z'; c++ { alphaNumeric[c] = true }
	i, j := 0, len(s) - 1
	s = strings.ToLower(s)
	for i < j {
		if !alphaNumeric[s[i]] {
			i++
			continue
		}
		if !alphaNumeric[s[j]] {
			j--
			continue
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
