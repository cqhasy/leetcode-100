package leetcode_100

func groupAnagrams(strs []string) [][]string {
	var res [][]string
	m := make(map[string]int)
	var x = 0
	if len(strs) == 0 {
		return [][]string{}

	}
	for k, v := range strs {
		if k == 0 {
			res = append(res, []string{v})
			m[sortStr(v)] = x
			x++
			continue
		}
		val, ok := m[sortStr(v)]
		if !ok {
			res = append(res, []string{v})
			m[sortStr(v)] = x
			x++
			continue
		}
		res[val] = append(res[val], v)

	}
	return res
}
func sortStr(s string) string {
	b := []byte(s)
	var m byte
	for i := 0; i < len(b); i++ {
		for j := i + 1; j < len(b); j++ {
			if b[i] > b[j] {
				m = b[i]
				b[i] = b[j]
				b[j] = m
			}
		}
	}
	return string(b)
}

//func groupAnagrams(strs []string) [][]string {
//	m := make(map[[26]byte][]string)
//	for i := range strs {
//		tmp := strs[i]
//		key := [26]byte{}
//		for _, v := range tmp {
//			key[v-'a']++
//		}
//		m[key] = append(m[key], tmp)
//	}
//	res := make([][]string, 0, len(m))
//	for _, value := range m {
//		res = append(res, value)
//	}
//	return res
//}
