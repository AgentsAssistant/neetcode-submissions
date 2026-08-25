func minWindow(s string, t string) string {
    if m := len(t); m == 0 || m > len(s) {
		return ""
	}
	need := make(map[byte]int)
	for k := range t {
		need[t[k]]++
	}
	have := make(map[byte]int)
	i, j := 0, 0
	res := ""
	for ; j < len(s); j++ {
		// fmt.Println(j)
		if !isValid(need, have){
			have[s[j]]++
		} 
		for isValid(need, have) {
			// fmt.Println(s[i:j+1], have, need)
			if res == "" || j-i+1 < len(res) {
				res = s[i:j+1]
			}
			have[s[i]]--
			i++
		}
	}
	return res
}

func isValid(need, have map[byte]int) bool {
	for ch, count := range need {
		if c, ok := have[ch]; !ok || c < count {
			return false
		}
	}
	return true
}
