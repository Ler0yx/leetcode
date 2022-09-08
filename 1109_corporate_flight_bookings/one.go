package corporateFlightBookings

import "fmt"

func CorpFlightBookings(bookings [][]int, n int) []int {
	flights := make(map[int]int)
	seatsBooked := make([]int, n)

	for i := 0; i <= n; i++ {
		flights[i] = 0
	}
	fmt.Println(flights)
	fmt.Println(seatsBooked)

	for i := 0; i < len(bookings); i++ {
		for j := bookings[i][0]; j <= bookings[i][1]; j++ {
			flights[j] += bookings[i][2]
		}
	}
	for i := 0; i < len(flights); i++ {
		seatsBooked[i] = flights[i]
	}
	return seatsBooked[1:]
}
