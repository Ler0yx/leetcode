package happyNumber

import (
	"fmt"
	"strconv"
	"strings"
)

func isHappy(n int) bool {

	sum := 0

	results := make(map[int]int)

	nString := fmt.Sprintf("%d", n)
	nSlice := strings.Split(nString, "")

	for {
		for i, v := range nSlice {
			val, _ := strconv.Atoi(v)
			sum += val * val
			fmt.Printf("i: %d, sum: %d\n", i, sum)
		}

		if sum == 1 {
			return true
		} else if _, exists := results[sum]; exists {
			return false
		}
		results[sum] = 1
		fmt.Printf("%v\n", results)
		nString = fmt.Sprintf("%d", sum)
		nSlice = strings.Split(nString, "")
		sum = 0
	}

}

func main() {
	fmt.Println(isHappy(2))
}
