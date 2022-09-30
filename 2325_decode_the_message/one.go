package decodeTheMessage

func decodeMessage(key string, message string) string {

	var alphabet = "abcdefghijklmnopqrstuvwxyz"
	substitution := make(map[byte]int)
	var result = make([]byte, len(message))
	var counter int

	for i := 0; i < len(key); i++ {
		if _, ok := substitution[key[i]]; !ok {
			if key[i] == 32 {
				continue
			} else {
				substitution[key[i]] = counter
				counter++
			}

		}
	}

	for i := 0; i < len(message); i++ {
		if message[i] == 32 {
			result[i] = 32
		} else {
			tmp := substitution[message[i]]
			result[i] = alphabet[tmp]
		}
	}

	return string(result)
}
