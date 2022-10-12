package robotwalk

var Position int
var North = 0
var East = 1
var South = 2
var West = 3

var Table [9][9]string

var X int
var Y int

func ReadWalk(walk string) []string {
	// var arr []string
	// for i, c := range walk {
	// 	arr[i] = string(c)
	// }
	return []string{"L", "W", "W", "W", "W", "L", "W"}

}

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
