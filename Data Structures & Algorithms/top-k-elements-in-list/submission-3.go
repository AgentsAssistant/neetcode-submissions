func topKFrequent(nums []int, k int) []int {
	count := map[int]int{}
	for _, n := range nums {
		count[n]++
	}
	bucket := make([][]int, len(nums)+1)
	for v, c := range count {
		bucket[c] = append(bucket[c], v)
	}
	res := []int{}
	i := len(bucket) - 1
	for i >= 0 {
		for _, n := range bucket[i] {
			res = append(res, n)
			if len(res) == k {
				return res
			}
		}
		i--
	}
	return res
}
