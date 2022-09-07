package floodFill

func floodFill2(image [][]int, sr int, sc int, color int) [][]int {
	func floodFill(image [][]int, sr int, sc int, color int) [][]int {
		startColor := image[sr][sc]
		targetColor := color
		
		if startColor != targetColor {
			
			image[sr][sc] = targetColor
			
			//Sets the pixel below image[sr][sc] to the target color, if it exists and has the starting color
			if sr-1 >= 0 {
				if image[sr-1][sc] == startColor {
					image[sr-1][sc] = targetColor
				}
	
				//Iterates the lower part of the image, starting from the starting row -1
				for i := sr-1; i >= 0; i-- {
	
					//Iterates the left part of the row, starting from "sc" -1
					for j := sc-1; j >= 0; j-- {
						if image[i][j+1] == targetColor && image[i][j] == startColor {
							image[i][j] = targetColor
						} else {
							break
						}
					}
	
					//Iterates the right part of the row, starting from "sc" +1
					for k := sc+1; k < len(image[i]); k++ {
						if image[i][k-1] == targetColor && image[i][k] ==startColor {
							image[i][k] = startColor
						} else {
							break
						}
					}
				}
			}
	
			//Sets the pixel above image[sr][sc] to the target color, if it exists and has the starting color
			if sr+1 < len(image){
				if image[sr+1][sc] == startColor {
					image[sr+1][sc] = targetColor
				}
	
				//Iterates the upper part of the image, starting from the starting row +1
				for i := sr+1; i < len(image); i++ {
	
					//Iterates the left part of the row, starting from "sc" -1
					for j := sc-1; j >= 0; j-- {
						if image[i][j+1] == targetColor && image[i][j] == startColor {
							image[i][j] = targetColor
						}
					}
	
					//Iterates the right part of the row, starting from "sc" +1
					for k := sc+1; k < len(image); k++ {
						if image[i][k-1] == targetColor && image[i][k] == startColor {
							image[i][k] = targetColor
						}
					}
				}
			}
		}
		return image
	}
}
