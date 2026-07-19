func groupAnagrams(strs []string) [][]string {
	countGroup := make(map[[26]int][]string)
	for _, s := range strs {
		var count [26]int
		for _, c := range s {
			count[c-'a']++
		}
		countGroup[count] = append(countGroup[count], s)
	}
	result := make([][]string, 0)
	for _, group := range countGroup {
		result = append(result, group)
	}
	return result
}
