package findTheHighestAltitude

func largestAltitude(gain []int) int {
    var max int
    var idx int
    var sum int
    idx += 1

    for i := 0; i < len(gain); i++ {
        sum += gain[i]
        if sum > max {
            max = sum
            idx = i+1
        }
    }
    return max
}