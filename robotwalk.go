package robotwalk

var Position int
var North = 0
var East = 1
var South = 2
var West = 3

var Table [9][9]string

var X int
var Y int

var Row int
var Column int

func SetStartWalking() {
	setStartXY()
	setStartTable()
	setStartPosition()
}

func setStartPosition() {
	Position = North
}

func setStartXY() {
	X = 0
	Y = 0
}

func setStartTable() {
	Row = 4
	Column = 4
	Table[Row][Column] = "0"
	for column := 0; column < 9; column++ {
		for row := 0; row < 9; row++ {
			Table[row][column] = "*"
		}
	}
}

// func RobotWalk(walk string) {
// 	for _, c := range walk {
// 		CalculateRobotWalk(string(c))
// 	}
// }

// func CalculateRobotWalk(walk string) {
// 	if walk != "W" {
// 		if walk == "L" {
// 			if Position == North {
// 				Position = West
// 			} else {
// 				Position--
// 			}
// 		} else {
// 			if Position == West {
// 				Position = North
// 			} else {
// 				Position++
// 			}
// 		}
// 	} else {
// 		if Position == North {
// 			Y++
// 		} else if Position == East {
// 			X++
// 		} else if Position == South {
// 			Y--
// 		} else if Position == West {
// 			X--
// 		}
// 	}
// }

// func setXY(x int, y int) {
// 	X = x
// 	Y = y
// }

// func SetResult(walkArray string) (int, int) {
// 	x := 0
// 	y := 0
// 	position := North
// 	setTable()

// 	for _, c := range walkArray {
// 		x, y, position = CalculateNextPosition(x, y, position, string(c))
// 		SetPositionRobotWalk(x, y)
// 	}
// 	return x, y
// }

// func CalculateNextPosition(x int, y int, position int, walk string) (int, int, int) {
// 	if walk != "W" {
// 		if walk == "L" {
// 			if position == North {
// 				position = West
// 			} else {
// 				position--
// 			}
// 		} else {
// 			if position == West {
// 				position = North
// 			} else {
// 				position++
// 			}
// 		}
// 	} else {
// 		if position == North {
// 			y++
// 		} else if position == East {
// 			x++
// 		} else if position == South {
// 			y--
// 		} else if position == West {
// 			x--
// 		}
// 	}
// 	return x, y, position
// }

// func SetPositionRobotWalk(x int, y int) {
// 	Table[y+4][x+4] = "0"
// }

// func setTable() {
// 	Table[4][4] = "0"
// 	for i := 0; i < 9; i++ {
// 		for j := 0; j < 9; j++ {
// 			Table[j][i] = "*"
// 		}
// 	}
// }
