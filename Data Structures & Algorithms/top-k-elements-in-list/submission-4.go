func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for _, n := range nums {
		count[n]++
	}
	bucket := make([][]int, len(nums)+1)
	for n, c := range count {
		bucket[c] = append(bucket[c], n)
	}
	var res []int
	for i := len(bucket)-1; k > 0; i-- {
		if l := len(bucket[i]); l > 0 {
			res = append(res, bucket[i]...)
			k -= l
		}
	}
	return res
}
