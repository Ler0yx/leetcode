package jewelsAndStones

func numJewelsInStones(jewels string, stones string) int {
    amounts := make(map[byte]int)
    var sum int

    for i := 0; i < len(stones); i++ {
        amounts[stones[i]] += 1
    }

    for j := 0; j < len(jewels); j++ {
        sum += amounts[jewels[j]]
    }
    return sum
}