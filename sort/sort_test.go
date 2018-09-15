package sort_test

func equals(nums, nums2 []int) bool {
	if len(nums) != len(nums2) {
		return false
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != nums2[i] {
			return false
		}
	}
	return true
}
