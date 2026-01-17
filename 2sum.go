package go_tips

func twoSum(nums []int, target int) []int {
	existingNums := make(map[int]int)

	for index, num := range nums {
		if value, ok := existingNums[target-num]; ok {
			return []int{value, index}
		}

		existingNums[num] = index
	}

	return nil
}
