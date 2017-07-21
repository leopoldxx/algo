package test

// SubArray2 xxx
func SubArray2(nums []int, min, max int) int {
	if len(nums) == 0 {
		return 0
	}
	prefix := make([]int, len(nums))
	prefix[0] = nums[0]

	for idx := 1; idx < len(nums); idx++ {
		prefix[idx] = prefix[idx-1] + nums[idx]
	}

	result := 0
	for i := 0; i < len(prefix); i++ {
		if prefix[i] >= min && prefix[i] <= max {
			result++
		}
		for j := i + 1; j < len(prefix); j++ {

			diff := prefix[j] - prefix[i]
			if diff >= min && diff <= max {
				result++
			}
		}
	}

	return result
}
