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
		myRobot.WalkingTable = SetWalkingTable(myRobot.WalkingTable, myRobot.Table)
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
	robot.Table = Table{Row: 4, Column: 4}
	robot.WalkingTable[robot.Table.Row][robot.Table.Column] = "0"
	return robot
}

func CalculateNextPosition(walk string, robot Robot) (Robot, error) {
	if walk != "W" {
		robot.Direction = ChangeDirection(walk, robot.Direction)
	} else {
		robot.Position, robot.Table = ChangePosition(robot.Direction, robot.Position, robot.Table)
		err := ValidateInMyTable(robot.Position)
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

func ChangePosition(direction Direction, position Position, table Table) (Position, Table) {
	if direction == North {
		position.Y++
		table.Row--
	} else if direction == East {
		position.X++
		table.Column++
	} else if direction == South {
		position.Y--
		table.Row++
	} else {
		position.X--
		table.Column--
	}
	return position, table
}

func SetWalkingTable(walkingTable [9][9]string, table Table) [9][9]string {
	walkingTable[table.Row][table.Column] = "0"
	return walkingTable
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

func ValidateInMyTable(position Position) error {
	if position.X > 4 || position.X < -4 || position.Y > 4 || position.Y < -4 {
		return errors.New("can't walk")
	}
	return nil
}
