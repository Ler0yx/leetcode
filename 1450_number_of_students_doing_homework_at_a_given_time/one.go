package numberOfStudentsDoingHomeworkAtAGivenTime

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	var counter int

	for i := 0; i < len(startTime); i++ {
		if startTime[i] <= queryTime && endTime[i] >= queryTime {
			counter += 1
		}
	}
	return counter
}
