package findTheWinnerOfAnArrayGame

func getWinner(arr []int, k int) int {
	var counter int

	for counter < k {
		if arr[0] > arr[1] {
			tmp := arr[1]
			arr = append(arr[0:1], arr[2:]...)
			arr = append(arr, tmp)
			counter++
		} else if arr[0] < arr[1] {
			tmp := arr[0]
			arr = append(arr[1:], tmp)
		}
	}
	return arr[0]
}
