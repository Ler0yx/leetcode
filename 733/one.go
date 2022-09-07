package floodFill

import "fmt"

func floodFill(image [][]int, sr int, sc int, color int) [][]int {

	startColor := image[sr][sc]
	targetColor := color
	image[sr][sc] = targetColor

	//Handles the lower part of the image
	for i := sr; i >= 0; i-- {

		if image[i+1][sc] == targetColor && image[i][sc] == startColor {
			image[i][sc] = targetColor
		} else if i != sr {
			break
		}

		//Iterates to the left (downwards), starting from start position
		for j := sc - 1; j >= 0; j-- {
			if image[i][j+1] == targetColor && image[i][j] == startColor {
				fmt.Println(i)
				fmt.Println(j)

				image[i][j] = targetColor
			} else {
				break
			}
		}
		//Iterates to the right (downwards), starting from start position
		for k := sc + 1; k < len(image[i]); k++ {
			if image[i][k-1] == targetColor && image[i][k] == startColor {
				image[i][k] = targetColor
			} else {
				break
			}
		}
	}
	//Handles the upper part of the image
	for i := sr + 1; i < len(image); i++ {

		if image[i-1][sc] == targetColor && image[i][sc] == startColor {
			image[i][sc] = targetColor
		} else {
			fmt.Println("i:", i)
			break
		}

		//Iterates to the right (upwards), starting from start position
		for j := sr - 1; j >= 0; i++ {
			if image[i][j+1] == targetColor && image[i][j] == startColor {
				image[i][j] = targetColor
			} else {
				break
			}
		}
		//Iterates to the right (upwards), starting from start position
		for k := sr + 1; k < len(image[i]); i++ {
			if image[i][k-1] == targetColor && image[i][k] == startColor {
				image[i][k] = targetColor
			} else {
				break
			}
		}
	}
	return image
}
