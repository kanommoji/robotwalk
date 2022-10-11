package robotwalk

func CalculateNextPosition(x int, y int, position int, walk string) (int, int, int) {
	if walk != "W" {
		if walk == "L" {
			if position == 0 {
				position = 3
			} else {
				position--
			}
		} else {
			if position == 3 {
				position = 0
			} else {
				position++
			}
		}
	} else {
		if position == 0 {
			y++
		} else if position == 1 {
			x++
		} else if position == 2 {
			y--
		} else if position == 3 {
			x--
		}
	}

	return x, y, position
}
