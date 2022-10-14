package robotwalk

import (
	"errors"
	"strconv"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type Position struct {
	X int
	Y int
}

type Table struct {
	Row    int
	Column int
}

type Robot struct {
	Direction    Direction
	WalkingTable [9][9]string
	Position     Position
	Table        Table
}

func RobotWalk(walks string) (string, error) {
	myRobot := InitRobot()
	var err error
	for _, walk := range walks {
		myRobot, err = CalculateNextPosition(string(walk), myRobot)
		if err != nil {
			return "", err
		}
		myRobot = setWalkingTable(myRobot)
	}
	return readResult(myRobot), err
}

func InitRobot() Robot {
	var robot Robot
	for column := 0; column < 9; column++ {
		for row := 0; row < 9; row++ {
			robot.WalkingTable[row][column] = "*"
		}
	}
	robot.Table.Row = 4
	robot.Table.Column = 4
	robot.WalkingTable[robot.Table.Row][robot.Table.Column] = "0"
	return robot
}

func CalculateNextPosition(walk string, robot Robot) (Robot, error) {
	if walk != "W" {
		robot.Direction = ChangeDirection(walk, robot.Direction)
	} else {
		robot.Position.X, robot.Position.Y, robot.Table.Row, robot.Table.Column = ChangePosition(robot.Direction, robot.Position.X, robot.Position.Y, robot.Table.Row, robot.Table.Column)
		err := ValidateInMyTable(robot.Position.X, robot.Position.Y)
		if err != nil {
			return Robot{}, err
		}
	}
	return robot, nil
}

func ChangeDirection(walk string, direction Direction) Direction {
	if walk == "L" {
		if direction == North {
			direction = West
		} else {
			direction--
		}
	} else {
		if direction == West {
			direction = North
		} else {
			direction++
		}
	}
	return direction
}

func ChangePosition(direction Direction, x int, y int, row int, column int) (int, int, int, int) {
	if direction == North {
		y++
		row--
	} else if direction == East {
		x++
		column++
	} else if direction == South {
		y--
		row++
	} else if direction == West {
		x--
		column--
	}
	return x, y, row, column
}

func setWalkingTable(robot Robot) Robot {
	robot.WalkingTable[robot.Table.Row][robot.Table.Column] = "0"
	return robot
}

func readResult(robot Robot) string {
	position := "Position : (" + strconv.Itoa(robot.Position.X) + "," + strconv.Itoa(robot.Position.Y) + ")"
	var table string
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			table += robot.WalkingTable[row][column]
		}
		table += "\n"
	}
	return position + "\n" + table
}

func ValidateInMyTable(x int, y int) error {
	if x > 4 || x < -4 || y > 4 || y < -4 {
		return errors.New("can't walk")
	}
	return nil
}
