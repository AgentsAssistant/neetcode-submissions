func twoSum(nums []int, target int) []int {
   seen := make(map[int]int)
   var vj int
   for i, n := range nums {
		vj = target - n
		if j, ok := seen[vj]; ok {
			return _sort(i, j)
		}
		seen[n] = i
   }
   return nil
}

func _sort(a, b int) []int {
	if a < b {
		return []int{a,b}
	}
	return []int{b,a}
}