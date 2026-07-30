func groupAnagrams(strs []string) [][]string {
	group := make(map[[26]int][]string)
	for _, s := range strs {
		var count [26]int
		for _, c := range s {
			count[c-'a']++
		}
		group[count] = append(group[count], s)
	}
	var res [][]string
	for _, v := range group {
		res = append(res, v)
	}
	return res
}
