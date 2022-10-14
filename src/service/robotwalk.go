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
		if robot.Direction == North {
			robot.Position.Y++
			robot.Table.Row--
		} else if robot.Direction == East {
			robot.Position.X++
			robot.Table.Column++
		} else if robot.Direction == South {
			robot.Position.Y--
			robot.Table.Row++
		} else if robot.Direction == West {
			robot.Position.X--
			robot.Table.Column--
		}
		err := validateInMyTable(robot.Position.X, robot.Position.Y)
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

func validateInMyTable(x int, y int) error {
	if x > 4 || x < -4 || y > 4 || y < -4 {
		return errors.New("can't walk")
	}
	return nil
}
