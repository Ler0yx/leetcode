package averageSalaryExcludingTheMinimumAndMaximumSalary

import (
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
