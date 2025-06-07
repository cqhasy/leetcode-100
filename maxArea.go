package leetcode_100

func maxArea(height []int) int {
	var h, w, re, t int
	if len(height) <= 1 {
		return 0
	}
	t = len(height)
	h = height[t-1]
	w = len(height) - 1
	re = 0
	for i := 0; i < len(height) && w > 0; {
		if height[i] <= h {
			if height[i]*w > re {
				re = height[i] * w
			}
			i++
			w--
			continue
		} else {
			if h*w > re {
				re = h * w
			}
			t--
			w--
			h = height[t-1]

		}

	}
	return re
}

//官方题解
//func maxArea(height []int) int {
//	p1 := 0
//	p2 := len(height) - 1
//	max := 0
//	for p1 < p2 {
//		var tmp int
//		if height[p1] < height[p2] {
//			tmp = (p2 - p1) * height[p1]
//			p1++
//		} else {
//			tmp = (p2 - p1) * height[p2]
//			p2--
//		}
//		if max < tmp {
//			max = tmp
//		}
//	}
//	return max
//}
