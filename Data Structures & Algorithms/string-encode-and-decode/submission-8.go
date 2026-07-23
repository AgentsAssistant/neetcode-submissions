type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var sb strings.Builder
	for _, s := range strs {
		sb.WriteString(strconv.Itoa(len(s)))
		sb.WriteRune('#')
		sb.WriteString(s)
	}
	return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
	if len(encoded) == 0 {
		return nil
	}
	var strs []string
	i := 0
	for j := 0; j < len(encoded); j++ {
		if encoded[j] == '#' {
			l, _ := strconv.Atoi(encoded[i:j])
			i = j+l+1
			strs = append(strs, encoded[j+1:i])
			j = i
		}
	}
	return strs
}
