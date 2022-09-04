package shiftingLetters

func shiftAmount(arr []int) []int {
    
    solution := make([]int, len(arr))
    
    for i := 0; i < len(arr); i++ {
        if arr[i] > 25 {
            arr[i] -= (arr[i] / 26) * 26
        }
        solution[i] = arr[i]
    }
    return solution
}

func stringToInt(s string) []int {
    
    solution := make([]int, len(s))
    
    for i := 0; i < len(s); i++ {
        solution[i] = int(s[i])
    }    
    return solution
}

func shiftingLetters4(s string, shifts []int) string {
    
    var intSlice = stringToInt(s)
    var shiftsFinal = shiftAmount(shifts)
    byteSlice := make([]byte, len(s))
	var tmp int

    for i := len(s) - 1; i >= 0; i-- {
		
		tmp += shiftsFinal[i]
        if tmp > 25 {
			tmp -= (tmp / 26) * 26
		}
        tmp2 := intSlice[i] + tmp
        if tmp2 > 122 {
            tmp2 -= 26
        }
        byteSlice[i] = byte(tmp2)
	    }    
    return string(byteSlice)
}

