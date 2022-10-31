package twoFurthestHousesWithDifferentColors

func abs (x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func maxDistance(colors []int) int {
    
    var dst int
    
    for i := 0; i < len(colors); i++ {
        for j := len(colors)-1; j > i; j-- {
            if colors[i] != colors[j] {
                if abs(i-j) > dst {
                    dst = abs(i-j)
                break
                }              
            }
        }
        if i+dst >= len(colors) {
            break
        }
    }
    return dst
}