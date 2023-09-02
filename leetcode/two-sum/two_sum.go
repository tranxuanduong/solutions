package two_sum

func twoSumBruceForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSumWithMap(nums []int, target int) []int {
	m := make(map[int]int)
	for i := range nums {
		if idx, ok := m[target-nums[i]]; ok {
			return []int{idx, i}
		} else {
			m[nums[i]] = i
		}
	}
	return nil
}
