type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0 {
		return "0x0"
	}
	return strings.Join(strs, "a1a")
}

func (s *Solution) Decode(encoded string) []string {
	if encoded == "0x0" {
		return nil
	}
	return strings.Split(encoded, "a1a")
}
