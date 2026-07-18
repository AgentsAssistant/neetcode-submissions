func twoSum(nums []int, target int) []int {
   seen := make(map[int]int)
   var vj int
   for i, n := range nums {
		vj = target - n
		if j, ok := seen[vj]; ok {
			return []int{j,i}
		}
		seen[n] = i
   }
   return nil
}