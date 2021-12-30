package leetcode

// 27.
func removeElement(nums []int, val int) int {
	n := len(nums)
	for i := 0; i < n; {
		if nums[i] == val {
			n--
			nums[i] = nums[n]
		} else {
			i++
		}
	}
	return n
}

// 2.
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		idx, ok := m[target-v]
		if ok {
			return []int{idx, k}
		}
		m[v] = k
	}
	return nil
}
