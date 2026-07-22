func groupAnagrams(strs []string) [][]string {
	groups := map[[26]int][]string{}
	for _, s := range strs {
		var count [26]int
		for _, c := range s {
			count[c-'a']++
		}
		groups[count] = append(groups[count], s)
	}
	res := [][]string{}
	for _, group := range groups {
		res = append(res, group)
	}
	return res
}
