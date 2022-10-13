package robotwalk

import (
	"strconv"
)

type Robot struct {
	Direction    int
	WalkingTable [9][9]string
	Position     Position
	Table        Table
}

type Position struct {
	X int
	Y int
}

type Table struct {
	Row    int
	Column int
}

var MyRobot Robot
var North = 0
var East = 1
var South = 2
var West = 3

func RobotWalk(walk string) string {
	InitWalking()
	for _, c := range walk {
		CalculateNextDirection(string(c))
		setTable(MyRobot.Table.Row, MyRobot.Table.Column)
	}
	return readResult()
}

func InitWalking() {
	initTable()
	initPosition()
	initDirection()
}

func initDirection() {
	MyRobot.Direction = North
}

func initPosition() {
	MyRobot.Position.X = 0
	MyRobot.Position.Y = 0
}

func initTable() {
	for column := 0; column < 9; column++ {
		for row := 0; row < 9; row++ {
			MyRobot.WalkingTable[row][column] = "*"
		}
	}
	MyRobot.Table.Row = 4
	MyRobot.Table.Column = 4
	MyRobot.WalkingTable[MyRobot.Table.Row][MyRobot.Table.Column] = "0"
}

func CalculateNextDirection(walk string) {
	if walk != "W" {
		if walk == "L" {
			if MyRobot.Direction == North {
				MyRobot.Direction = West
			} else {
				MyRobot.Direction--
			}
		} else {
			if MyRobot.Direction == West {
				MyRobot.Direction = North
			} else {
				MyRobot.Direction++
			}
		}
	} else {
		if MyRobot.Direction == North && MyRobot.Position.Y < 4 {
			MyRobot.Position.Y++
			MyRobot.Table.Row--
		} else if MyRobot.Direction == East && MyRobot.Position.X < 4 {
			MyRobot.Position.X++
			MyRobot.Table.Column++
		} else if MyRobot.Direction == South && MyRobot.Position.Y > -4 {
			MyRobot.Position.Y--
			MyRobot.Table.Row++
		} else if MyRobot.Direction == West && MyRobot.Position.X > -4 {
			MyRobot.Position.X--
			MyRobot.Table.Column--
		}

	}
}

func setTable(row int, column int) {
	MyRobot.WalkingTable[row][column] = "0"
}

func readResult() string {
	return readXYPoint() + "\n" + readTable()
}

func readXYPoint() string {
	return "Position : (" + strconv.Itoa(MyRobot.Position.X) + "," + strconv.Itoa(MyRobot.Position.Y) + ")"
}

func readTable() string {
	var table string
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			table += MyRobot.WalkingTable[row][column]
		}
		table += "\n"
	}
	return table
}
