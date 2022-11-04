package canPlaceFlowers

func canPlaceFlowers(flowerbed []int, n int) bool {
    
    var counter int
    
    for i, plot := range flowerbed {
        if plot == 1 {
            i++
        } else if plot == 0 && i < len(flowerbed)-2 && flowerbed[i+1] != 1 {
            counter += 1
        } else if plot == 0 && i == len(flowerbed)-1 {
            counter += 1
        }
    }
    if counter >= n {
        return true
    }
    return false
}