package checkIfItIsAStraightLine

func checkStraightLine(coordinates [][]int) bool {

	var m float64
	var n float64
	var x1 float64
	var y1 float64
	var x2 float64
	var xKonst float64
	var yKonst float64
	var y2 float64
	var x float64
	var y float64

	if len(coordinates) == 2 {
		return true
	}

	for i, coordinate := range coordinates {
		if i == 0 {
			for i, point := range coordinate {
				if i == 0 {
					x1 = float64(point)
				} else {
					y1 = float64(point)
				}
			}
		} else if i == 1 {
			for i, point := range coordinate {
				if i == 0 {
					x2 = float64(point)
				} else {
					y2 = float64(point)
				}
			}
		} else {
			break
		}
	}

	//Edge Case: vertical line
	if x2-x1 == 0 {
		xKonst = float64(x1)

		for _, coordinate := range coordinates {
			for i, point := range coordinate {
				if i == 0 {
					x = float64(point)
				}
				if x != xKonst {
					return false
				} else {
					continue
				}
			}
		}

		//Edge Case: horizontal line
	} else if y2-y1 == 0 {
		yKonst = float64(y1)

		for _, coordinate := range coordinates {
			for i, point := range coordinate {
				if i == 1 {
					y = float64(point)
				} else {
					continue
				}
				if y != yKonst {
					return false
				} else {
					continue
				}
			}
		}
	} else {
		m = (float64(y2) - float64(y1)) / (float64(x2) - float64(x1))
		n = ((y1 * x2) - (y2 * x1)) / (x2 - x1)

		for i, coordinate := range coordinates {
			if i > 1 {
				for i, point := range coordinate {
					if i == 0 {
						x = float64(point)
					} else {
						y = float64(point)
					}
				}
				if y == m*x+n {
					continue
				} else {
					return false
				}
			} else {
				continue
			}
		}
	}
	return true
}

func main() {

}
