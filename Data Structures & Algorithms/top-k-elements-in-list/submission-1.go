func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for _, n := range nums {
		count[n]++
	}
	bucket := make([][]int, len(nums))
	for n, freq := range count {
		bucket[freq-1] = append(bucket[freq-1], n)
	}
	var result []int
	n := len(bucket)-1
	for ; k > 0; n-- {
		if l := len(bucket[n]); l > 0 {
			result = append(result, bucket[n]...)
			k -= l
		}
	}
	return result
}
