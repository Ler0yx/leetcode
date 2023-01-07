package selfDividingNumbers

func selfDividingNumbers(left int, right int) []int {
	res := make([]int, 0)
	var num int
	for i := left; i <= right; i++ {
		tmp := 0
		num = i
		for num > 0 {
			tmp = num % 10
			num /= 10
			if tmp == 0 || i%tmp != 0 {
				break
			} else if num == 0 {
				res = append(res, i)
			}
		}
	}
	return res
}
