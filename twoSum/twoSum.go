package twoSum

func twoSum(nums []int, target int) []int {
	//use hashmap to trade storage for time
	var m = make(map[int]int)
	//only loop once, use hashmap to find compliment
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
