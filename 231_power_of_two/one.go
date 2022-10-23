package powerOfTwo

import "math"

func isPowerOfTwo(n int) bool {

	powTwo := make(map[int]bool)

	for i := 0; i < 31; i++ {
		powTwo[int(math.Pow(2, float64(i)))] = true
	}

	if _, ok := powTwo[n]; ok {
		return true
	}
	return false
}
