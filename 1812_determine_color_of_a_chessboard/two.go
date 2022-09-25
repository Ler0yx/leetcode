package determineColorOfAChessboard

import "fmt"

func SquareIsWhite(coordinates string) bool {
	row, col := coordinates[1]-'0', coordinates[0]-'a'+1

	fmt.Println((12)&1 == 1)
	return (row+col)&1 == 1
}
