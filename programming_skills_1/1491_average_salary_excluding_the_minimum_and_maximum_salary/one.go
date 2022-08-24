package main

import (
	"fmt"
	"sort"
)

func average(salary []int) float64 {

	var tmp float64

	sort.Ints(salary)
	for _, v := range salary[1 : len(salary)-1] {
		tmp += float64(v)
	}
	return tmp / (float64(len(salary)) - 2)
}

func main() {

	salary := []int{1000, 2000, 3000, 999999, 999998}

	fmt.Println(average(salary))

}
