package toLowerCase

import "fmt"

func toLowerCase(s string) string {
	byteSlice := []byte(s)
	for i := 0; i < len(byteSlice); i++ {

		if byteSlice[i] >= 65 && byteSlice[i] <= 90 {
			byteSlice[i] += 32
		}
	}
	return string(byteSlice)
}

func main() {
	s1 := "heLLo"

	fmt.Println(toLowerCase(s1))
}
