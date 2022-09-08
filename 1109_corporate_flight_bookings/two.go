package corporateFlightBookings

func CorpFlightBookings2(bookings [][]int, n int) []int {
	seatsBooked := make([]int, n)

	for i := 0; i < len(bookings); i++ {
		for j := bookings[i][0]; j <= bookings[i][1]; j++ {
			seatsBooked[j-1] += bookings[i][2]
		}
	}
	return seatsBooked
}
