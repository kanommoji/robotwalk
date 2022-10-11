package robotwalk

var North = 0
var East = 1
var South = 2
var West = 3

func CalculateNextPosition(x int, y int, position int, walk string) (int, int, int) {
	if walk != "W" {
		if walk == "L" {
			if position == North {
				position = West
			} else {
				position--
			}
		} else {
			if position == West {
				position = North
			} else {
				position++
			}
		}
	} else {
		if position == North {
			y++
		} else if position == East {
			x++
		} else if position == South {
			y--
		} else if position == West {
			x--
		}
	}
	return x, y, position
}
