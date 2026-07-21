type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var sb strings.Builder
	for _, s := range strs {
		sb.WriteString(strconv.Itoa(len(s)))
		sb.WriteString("#")
		sb.WriteString(s)
	}
	// 0#1#a5#hello
	return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
	k := len(encoded)
	if k == 0 {
		return nil
	}
	i := 0
	res := []string{}
	for j := 0; j < k; j++ {
		if encoded[j] == '#' {
			l, _ := strconv.Atoi(encoded[i:j])
			res = append(res, encoded[j+1:j+1+l])
			i = j+1+l
			j = i
		}
	}
	return res
}
