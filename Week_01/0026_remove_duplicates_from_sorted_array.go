package solutions

// use processedIndex to record processed last value, should be O(n)
func removeDuplicates(nums []int) int {

	if len(nums) == 1 {
		return 1
	}

	processedIndex := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[processedIndex] {
			processedIndex++
			nums[processedIndex] = nums[i]
		}
	}

	return processedIndex + 1
}
