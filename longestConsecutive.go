package leetcode_100

// 复杂度On
func longestConsecutive(nums []int) int {
	var re int = 0
	var cnt int = 0
	m := make(map[int]int)
	for _, v := range nums {
		m[v] = 1
	}

	for v, _ := range m {
		row := v
		for row < cnt || row > cnt+re {
			v++
			if _, ok := m[v]; ok {
				m[row]++
			} else {
				break
			}
		}
		if m[row] > re {
			re = m[row]
			cnt = row
		}
	}
	return re
}

//下面是官方题解，复杂度Onlogn
//func longestConsecutive(nums []int) int {
//	if len(nums) == 0 {
//		return 0
//	}
//	if len(nums) == 1 {
//		return 1
//	}
//	// 排序
//	sort.Ints(nums)
//
//	longest := 0
//	cur := 1
//	for i := 1; i < len(nums); i++ {
//		if nums[i] == nums[i-1]+1 {
//			cur++
//		} else if nums[i] == nums[i-1] {
//
//		} else {
//			cur = 1
//		}
//
//		if cur > longest {
//			longest = cur
//		}
//	}
//
//	return longest
//}
