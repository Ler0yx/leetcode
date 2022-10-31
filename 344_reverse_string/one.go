package reverseString

func reverseString(s []byte) []byte {

	last := len(s) / 2
	var swap byte
	i := 0
	j := len(s) - 1

	for i < last {
		swap = s[i]
		s[i] = s[j]
		s[j] = swap
		i++
		j--
	}
	return s
}
