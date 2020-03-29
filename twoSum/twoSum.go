package twoSum

func twoSum(nums []int, target int) []int {
	var m = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		compliment := target - nums[i]
		j, ok := m[compliment]
		if ok {
			return []int{i, j}
		}
		m[nums[i]] = i
	}
	return []int{}
}
