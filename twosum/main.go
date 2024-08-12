func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for index, val := range nums {
		complement := target - val
		if j, ok := seen[complement]; ok {
			return []int{j, index}
		}
		seen[val] = index
	}
	return nil
}
