package minimumAbsoluteDifference

func abs (x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func minimumAbsDifference(arr []int) [][]int {
    
    var minDiff := math.MaxInt()
    
    for i:=0; i < len(arr)-1; i++ {
        if abs(i) - abs(i+1) < minDiff {
            minDiff = 
        }
    }
}

1: den niedrigsten absoluten Abstand finden
2: aufsteigend sortieren
3: abstaende mit minimum vergleichen
4: paare sortieren