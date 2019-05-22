package ops

// Unique values of a slice 切片去重 - 和js去重方法中的一种类似
func Unique(s []float64) []float64 {
	m := map[float64]bool{}
	var r []float64

	for _, v := range s {
		if _, seen := m[v]; !seen {
			r = append(r, v)
			m[v] = true
		}
	}

	return r
}
