package util

func Max(s []float64, m float64) []float64 {
	maxSlice := make([]float64, len(s))
	for idx, v := range s {
		if m > v {
			maxSlice[idx] = m
		} else {
			maxSlice[idx] = v
		}
	}
	return maxSlice
}

func Sum(s []float64) float64 {
	sum := float64(0)
	for _, v := range s {
		sum += v
	}
	return sum
}

func NVals(val float64, n int) []float64 {
	vals := make([]float64, n)
	for idx, _ := range vals {
		vals[idx] = val
	}
	return vals
}

func DivideBy(s []float64, denominator float64) {
	for idx, x := range s {
		s[idx] = x / denominator
	}
}

func Clamp(s []float64, clampFunc func(float64) bool, clampVal float64) {
	for idx, x := range s {
		if clampFunc(x) {
			s[idx] = clampVal
		}
	}
}

func Map(s []float64, f func(float64) float64) []float64 {
	copy := make([]float64, len(s))
	for idx, x := range s {
		copy[idx] = f(x)
	}
	return copy
}

func MapInts(s []int, f func(int) int) []int {
	copy := make([]int, len(s))
	for idx, x := range s {
		copy[idx] = f(x)
	}
	return copy
}

func CopyInts(s []int) []int {
	return MapInts(s, func(x int) int { return x })
}

func Copy(s []float64) []float64 {
	return Map(s, func(x float64) float64 { return x })
}

func AddVectorTo(src []float64, dest []float64) {
	for idx, srcval := range src {
		dest[idx] += srcval
	}
}

func AddTo(srcval float64, dest []float64) {
	for idx := range dest {
		dest[idx] += srcval
	}
}

func MultBy(srcval float64, dest []float64) {
	for idx := range dest {
		dest[idx] *= srcval
	}
}

func Range(start, stop, inc int) []int {
	r := make([]int, (stop-start)/inc)
	for idx, _ := range r {
		r[idx] = start + idx*inc
	}
	return r
}
