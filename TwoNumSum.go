package leetcode_100

//	func twoSum(nums []int, target int) []int {
//		l := 0
//		r := len(nums) - 1
//
//		var res []int
//		for ; l < r; l++ {
//			t := target - nums[l]
//			for nums[r] != t && r > l {
//				r--
//			}
//			if nums[r] == t && r > l {
//
//				res = append(res, l)
//				res = append(res, r)
//				return res
//			}
//			r = len(nums) - 1
//		}
//		return nil
//
// }
func twoSum(nums []int, target int) []int {
	var m = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		t := target - nums[i]
		if j, ok := m[t]; ok {
			return []int{j, i}
		}
		m[nums[i]] = i
	}
	return nil
}
