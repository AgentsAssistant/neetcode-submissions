func isPalindrome(s string) bool {
	alphaNum := [128]bool{}
	for c := '0'; c <= '9'; c++ { alphaNum[c] = true }
	for c := 'a'; c <= 'z'; c++ { alphaNum[c] = true }
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	for i < j {
		if !alphaNum[s[i]] {
			i++
			continue
		}
		if !alphaNum[s[j]] {
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
