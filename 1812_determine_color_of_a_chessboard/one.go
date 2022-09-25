package determineColorOfAChessboard

func squareIsWhite(coordinates string) bool {
	switch {

	case coordinates[0]%2 == 0:
		if coordinates[1]%2 != 0 {
			return true
		}
		return false

	case coordinates[0]%2 != 0:
		if coordinates[1]%2 == 0 {
			return true
		}
		return false
	}
	return true
}
