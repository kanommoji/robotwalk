package robotwalk

import (
	"strconv"
)

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

func RobotWalk(walk string) {
	SetStartWalking()
	for _, c := range walk {
		CalculateNextPosition(string(c))
		setXY(X, Y)
		setTable(Row, Column)
	}
}

func SetStartWalking() {
	setStartTable()
	setStartXY()
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
	for column := 0; column < 9; column++ {
		for row := 0; row < 9; row++ {
			Table[row][column] = "*"
		}
	}
	Row = 4
	Column = 4
	Table[Row][Column] = "0"
}

func CalculateNextPosition(walk string) {
	if walk != "W" {
		if walk == "L" {
			if Position == North {
				Position = West
			} else {
				Position--
			}
		} else {
			if Position == West {
				Position = North
			} else {
				Position++
			}
		}
	} else {
		if Position == North && Y < 4 {
			Y++
			Row--
		} else if Position == East && X < 4 {
			X++
			Column++
		} else if Position == South && Y > -4 {
			Y--
			Row++
		} else if Position == West && X > -4 {
			X--
			Column--
		}

	}
}

func setXY(x int, y int) {
	X = x
	Y = y
}

func setTable(row int, column int) {
	Row = row
	Column = column
	Table[row][column] = "0"
}

func ReadResult() string {
	return readXYPoint() + "\n" + readTable()
}

func readXYPoint() string {
	return "Position : (" + strconv.Itoa(X) + "," + strconv.Itoa(Y) + ")"
}

func readTable() string {
	var table string
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			table += Table[row][column]
		}
		table += "\n"
	}
	return table
}
