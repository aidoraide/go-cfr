package util

func Max(s []float32, m float32) []float32 {
	maxSlice := make([]float32, len(s))
	for idx, v := range s {
		if m > v {
			maxSlice[idx] = m
		} else {
			maxSlice[idx] = v
		}
	}
	return maxSlice
}

func Sum(s []float32) float32 {
	sum := float32(0)
	for _, v := range s {
		sum += v
	}
	return sum
}

func NVals(val float32, n int) []float32 {
	vals := make([]float32, n)
	for idx, _ := range vals {
		vals[idx] = val
	}
	return vals
}

func DivideBy(s []float32, denominator float32) {
	for idx, x := range s {
		s[idx] = x / denominator
	}
}

func Clamp(s []float32, clampFunc func(float32) bool, clampVal float32) {
	for idx, x := range s {
		if clampFunc(x) {
			s[idx] = clampVal
		}
	}
}

func Map(s []float32, f func(float32) float32) []float32 {
	copy := make([]float32, len(s))
	for idx, x := range s {
		copy[idx] = f(x)
	}
	return copy
}

func Copy(s []float32) []float32 {
	return Map(s, func(x float32) float32 { return x })
}

func AddVectorTo(src []float32, dest []float32) {
	for idx, srcval := range src {
		dest[idx] += srcval
	}
}

func AddTo(srcval float32, dest []float32) {
	for idx := range dest {
		dest[idx] += srcval
	}
}

func MultBy(srcval float32, dest []float32) {
	for idx := range dest {
		dest[idx] *= srcval
	}
}
