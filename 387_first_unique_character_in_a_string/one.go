package firstUniqueCharacterInAString

func firstUniqChar(s string) int {
	d := -1
	counter := 0

	for i, letter := range s {
		counter = 0
		x := byte(letter)
		for d := 0; d <= len(s); d++ {
			if d == len(s) {
				return i
			} else {

				if s[d] == x {
					counter++
					if counter > 1 {
						break
					}
				}

			}
		}
	}
	return d
}
