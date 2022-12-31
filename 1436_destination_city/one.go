package destinationCity

func destCity(paths [][]string) string {
    cities := make(map[string]bool)
    for i := 0; i < len(paths); i++ {
            cities[paths[i][0]] = true
    }
    for j := 0; j < len(paths); j++ {
        if !cities[paths[j][1]] {
            return paths[j][1]
        }
    }
    return ""
}