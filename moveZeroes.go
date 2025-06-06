package leetcode_100

func moveZeroes(nums []int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}
	var l, r int = 0, 0
	for ; l <= len(nums); l++ {
		if l == len(nums) {
			for r < len(nums) {
				nums[r] = 0
				r++
			}
			break
		}
		if nums[l] != 0 {
			nums[r] = nums[l]
			r++
		}

	}

}
