func topKFrequent(nums []int, k int) []int {
	count := map[int]int{}
	for _, num := range nums {
		count[num]++
	}
	bucket := make([][]int, len(nums)+1)
	for num, q := range count {
		bucket[q] = append(bucket[q], num)
	}
	result := []int{}
	for i := len(bucket)-1; k > 0; i-- {
		if l := len(bucket[i]); l > 0 {
			result = append(result, bucket[i]...)
			k -= l
		}
	}
	return result
}
