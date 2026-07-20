type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var sb strings.Builder
	for _, s := range strs {
		sb.WriteString(strconv.Itoa(len(s)))
		sb.WriteString("#")
		sb.WriteString(s)
	}
	return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
	fmt.Println(encoded)
	if encoded == "" {
		return nil
	}
	n := len(encoded)
	strs := []string{}
	i := 0
	for j := 0; j < n; j++ {
		if encoded[j] == '#' {
			l, _ := strconv.Atoi(encoded[i:j])
			if l == 0 {
				strs = append(strs, "")
			} else {
				strs = append(strs, encoded[j+1:j+1+l])
			}
			i = j+1+l
			j = i
		}
	}
	return strs
}
