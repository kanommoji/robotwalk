package robotwalk_test

import (
	"errors"
	robotwalk "robotwalk/src/service"
	"testing"
)

func Test_RobotWalk_WWLWRWW(t *testing.T) {
	walk := "WWLWRWW"
	expected := `Position : (-1,4)
***0*****
***0*****
***00****
****0****
****0****
*********
*********
*********
*********
`
	actual, _ := robotwalk.RobotWalk(walk)

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_RobotWalk_LWWWWLW(t *testing.T) {
	walk := "LWWWWLW"
	expected := `Position : (-4,-1)
*********
*********
*********
*********
00000****
0********
*********
*********
*********
`
	actual, _ := robotwalk.RobotWalk(walk)

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_RobotWalk_LWWWWWWWWW_OutOfTable(t *testing.T) {
	walk := "LWWWWWWWWW"
	expectedErr := errors.New("can't walk")

	_, err := robotwalk.RobotWalk(walk)

	if expectedErr == err {
		t.Errorf("Expect %v but got %v", expectedErr, err)
	}
}

func Test_InitRobot(t *testing.T) {
	expected := robotwalk.Robot{
		Direction: robotwalk.North,
		WalkingTable: [9][9]string{
			{"*", "*", "*", "*", "*", "*", "*", "*", "*"},
			{"*", "*", "*", "*", "*", "*", "*", "*", "*"},
			{"*", "*", "*", "*", "*", "*", "*", "*", "*"},
			{"*", "*", "*", "*", "*", "*", "*", "*", "*"},
			{"*", "*", "*", "*", "0", "*", "*", "*", "*"},
			{"*", "*", "*", "*", "*", "*", "*", "*", "*"},
			{"*", "*", "*", "*", "*", "*", "*", "*", "*"},
			{"*", "*", "*", "*", "*", "*", "*", "*", "*"},
			{"*", "*", "*", "*", "*", "*", "*", "*", "*"}},
		Position: robotwalk.Position{X: 0, Y: 0},
		Table:    robotwalk.Table{Row: 4, Column: 4},
	}

	actual := robotwalk.InitRobot()

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_CalculateNextPosition_Walk_W_Direction_North_OutOfTable(t *testing.T) {
	walk := "W"
	robot := robotwalk.Robot{
		Direction: robotwalk.North,
		Position:  robotwalk.Position{X: 0, Y: 4},
		Table:     robotwalk.Table{Row: 0, Column: 4},
	}
	expected := robotwalk.Robot{
		Direction: robotwalk.North,
		Position:  robotwalk.Position{X: 0, Y: 0},
		Table:     robotwalk.Table{Row: 0, Column: 0},
	}

	actual, _ := robotwalk.CalculateNextPosition(walk, robot)

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_CalculateNextPosition_Walk_W_Direction_North(t *testing.T) {
	walk := "W"
	robot := robotwalk.Robot{
		Direction: robotwalk.North,
		Position:  robotwalk.Position{X: 0, Y: 0},
		Table:     robotwalk.Table{Row: 4, Column: 4},
	}
	expected := robotwalk.Robot{
		Direction: robotwalk.North,
		Position:  robotwalk.Position{X: 0, Y: 1},
		Table:     robotwalk.Table{Row: 3, Column: 4},
	}

	actual, _ := robotwalk.CalculateNextPosition(walk, robot)

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_CalculateNextPosition_Walk_W_Direction_East(t *testing.T) {
	walk := "W"
	robot := robotwalk.Robot{
		Direction: robotwalk.East,
		Position:  robotwalk.Position{X: 0, Y: 0},
		Table:     robotwalk.Table{Row: 4, Column: 4},
	}
	expected := robotwalk.Robot{
		Direction: robotwalk.East,
		Position:  robotwalk.Position{X: 1, Y: 0},
		Table:     robotwalk.Table{Row: 4, Column: 5},
	}

	actual, _ := robotwalk.CalculateNextPosition(walk, robot)

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_CalculateNextPosition_Walk_W_Direction_South(t *testing.T) {
	walk := "W"
	robot := robotwalk.Robot{
		Direction: robotwalk.South,
		Position:  robotwalk.Position{X: 0, Y: 0},
		Table:     robotwalk.Table{Row: 4, Column: 4},
	}
	expected := robotwalk.Robot{
		Direction: robotwalk.South,
		Position:  robotwalk.Position{X: 0, Y: -1},
		Table:     robotwalk.Table{Row: 5, Column: 4},
	}

	actual, _ := robotwalk.CalculateNextPosition(walk, robot)

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_CalculateNextPosition_Walk_W_Direction_West(t *testing.T) {
	walk := "W"
	robot := robotwalk.Robot{
		Direction: robotwalk.West,
		Position:  robotwalk.Position{X: 0, Y: 0},
		Table:     robotwalk.Table{Row: 4, Column: 4},
	}
	expected := robotwalk.Robot{
		Direction: robotwalk.West,
		Position:  robotwalk.Position{X: -1, Y: 0},
		Table:     robotwalk.Table{Row: 4, Column: 3},
	}

	actual, _ := robotwalk.CalculateNextPosition(walk, robot)

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_ChangeDirection_Walk_L_Change_Direction_North_To_West(t *testing.T) {
	walk := "L"
	direction := robotwalk.North
	expected := robotwalk.West

	actual := robotwalk.ChangeDirection(walk, direction)

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_ChangeDirection_Walk_L_Change_Direction_West_To_South(t *testing.T) {
	walk := "L"
	direction := robotwalk.West
	expected := robotwalk.South

	actual := robotwalk.ChangeDirection(walk, direction)

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_ChangeDirection_Walk_L_Change_Direction_South_North_To_East(t *testing.T) {
	walk := "L"
	direction := robotwalk.South
	expected := robotwalk.East

	actual := robotwalk.ChangeDirection(walk, direction)

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_ChangeDirection_Walk_L_Change_Direction_East_To_North(t *testing.T) {
	walk := "L"
	direction := robotwalk.East
	expected := robotwalk.North

	actual := robotwalk.ChangeDirection(walk, direction)

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_ChangeDirection_Walk_R_Change_Direction_North_To_East(t *testing.T) {
	walk := "R"
	direction := robotwalk.North
	expected := robotwalk.East

	actual := robotwalk.ChangeDirection(walk, direction)

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_ChangeDirection_Walk_R_Change_Direction_East_To_South(t *testing.T) {
	walk := "R"
	direction := robotwalk.East
	expected := robotwalk.South

	actual := robotwalk.ChangeDirection(walk, direction)

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_ChangeDirection_Walk_R_Change_Direction_South_To_West(t *testing.T) {
	walk := "R"
	direction := robotwalk.South
	expected := robotwalk.West

	actual := robotwalk.ChangeDirection(walk, direction)

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_ChangeDirection_Walk_R_Change_Direction_West_To_North(t *testing.T) {
	walk := "R"
	direction := robotwalk.West
	expected := robotwalk.North

	actual := robotwalk.ChangeDirection(walk, direction)

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_ChangePosition_Direction_North(t *testing.T) {
	direction := robotwalk.North
	position := robotwalk.Position{X: 0, Y: 0}
	table := robotwalk.Table{Row: 4, Column: 4}
	expectedPosition := robotwalk.Position{X: 0, Y: 1}
	expectedTable := robotwalk.Table{Row: 3, Column: 4}

	actualPosition, actualTable := robotwalk.ChangePosition(direction, position, table)

	if expectedPosition != actualPosition {
		t.Errorf("Expect X %v but got %v", expectedPosition, actualPosition)
	}
	if expectedTable != actualTable {
		t.Errorf("Expect Y %v but got %v", expectedTable, actualTable)
	}
}

func Test_ChangePosition_Direction_South(t *testing.T) {
	direction := robotwalk.South
	position := robotwalk.Position{X: 0, Y: 0}
	table := robotwalk.Table{Row: 4, Column: 4}
	expectedPosition := robotwalk.Position{X: 0, Y: -1}
	expectedTable := robotwalk.Table{Row: 5, Column: 4}

	actualPosition, actualTable := robotwalk.ChangePosition(direction, position, table)

	if expectedPosition != actualPosition {
		t.Errorf("Expect X %v but got %v", expectedPosition, actualPosition)
	}
	if expectedTable != actualTable {
		t.Errorf("Expect Y %v but got %v", expectedTable, actualTable)
	}
}

func Test_ChangePosition_Direction_East(t *testing.T) {
	direction := robotwalk.East
	position := robotwalk.Position{X: 0, Y: 0}
	table := robotwalk.Table{Row: 4, Column: 4}
	expectedPosition := robotwalk.Position{X: 1, Y: 0}
	expectedTable := robotwalk.Table{Row: 4, Column: 5}

	actualPosition, actualTable := robotwalk.ChangePosition(direction, position, table)

	if expectedPosition != actualPosition {
		t.Errorf("Expect X %v but got %v", expectedPosition, actualPosition)
	}
	if expectedTable != actualTable {
		t.Errorf("Expect Y %v but got %v", expectedTable, actualTable)
	}
}

func Test_ChangePosition_Direction_West(t *testing.T) {
	direction := robotwalk.West
	position := robotwalk.Position{X: 0, Y: 0}
	table := robotwalk.Table{Row: 4, Column: 4}
	expectedPosition := robotwalk.Position{X: -1, Y: 0}
	expectedTable := robotwalk.Table{Row: 4, Column: 3}

	actualPosition, actualTable := robotwalk.ChangePosition(direction, position, table)

	if expectedPosition != actualPosition {
		t.Errorf("Expect X %v but got %v", expectedPosition, actualPosition)
	}
	if expectedTable != actualTable {
		t.Errorf("Expect Y %v but got %v", expectedTable, actualTable)
	}
}

func Test_SetWalkingTable(t *testing.T) {
	walkingTable := [9][9]string{
		{"*", "*", "*", "*", "*", "*", "*", "*", "*"},
		{"*", "*", "*", "*", "*", "*", "*", "*", "*"},
		{"*", "*", "*", "*", "*", "*", "*", "*", "*"},
		{"*", "*", "*", "*", "*", "*", "*", "*", "*"},
		{"*", "*", "*", "*", "0", "*", "*", "*", "*"},
		{"*", "*", "*", "*", "*", "*", "*", "*", "*"},
		{"*", "*", "*", "*", "*", "*", "*", "*", "*"},
		{"*", "*", "*", "*", "*", "*", "*", "*", "*"},
		{"*", "*", "*", "*", "*", "*", "*", "*", "*"}}
	table := robotwalk.Table{Row: 0, Column: 0}
	expected := [9][9]string{
		{"0", "*", "*", "*", "*", "*", "*", "*", "*"},
		{"*", "*", "*", "*", "*", "*", "*", "*", "*"},
		{"*", "*", "*", "*", "*", "*", "*", "*", "*"},
		{"*", "*", "*", "*", "*", "*", "*", "*", "*"},
		{"*", "*", "*", "*", "0", "*", "*", "*", "*"},
		{"*", "*", "*", "*", "*", "*", "*", "*", "*"},
		{"*", "*", "*", "*", "*", "*", "*", "*", "*"},
		{"*", "*", "*", "*", "*", "*", "*", "*", "*"},
		{"*", "*", "*", "*", "*", "*", "*", "*", "*"}}

	actual := robotwalk.SetWalkingTable(walkingTable, table)

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_ValidateInMyTable_X_5(t *testing.T) {
	position := robotwalk.Position{X: 5, Y: 0}
	expected := errors.New("can't walk")

	actual := robotwalk.ValidateInMyTable(position)

	if actual == nil {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_ValidateInMyTable_X_minus_5(t *testing.T) {
	position := robotwalk.Position{X: -5, Y: 0}
	expected := errors.New("can't walk")

	actual := robotwalk.ValidateInMyTable(position)

	if actual == nil {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_ValidateInMyTable_Y_5(t *testing.T) {
	position := robotwalk.Position{X: 0, Y: 5}
	expected := errors.New("can't walk")

	actual := robotwalk.ValidateInMyTable(position)

	if actual == nil {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_ValidateInMyTable_Y_minus_5(t *testing.T) {
	position := robotwalk.Position{X: 0, Y: -5}
	expected := errors.New("can't walk")

	actual := robotwalk.ValidateInMyTable(position)

	if actual == nil {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_ValidateInMyTable_X_4_Y_minus_5(t *testing.T) {
	position := robotwalk.Position{X: 4, Y: -5}
	expected := errors.New("can't walk")

	actual := robotwalk.ValidateInMyTable(position)

	if actual == nil {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_ValidateInMyTable_X_4(t *testing.T) {
	position := robotwalk.Position{X: 4, Y: 0}
	actual := robotwalk.ValidateInMyTable(position)

	if actual != nil {
		t.Errorf("Expect nil but got %v", actual)
	}
}

func Test_ValidateInMyTable_X_minus_4(t *testing.T) {
	position := robotwalk.Position{X: -4, Y: 0}
	actual := robotwalk.ValidateInMyTable(position)

	if actual != nil {
		t.Errorf("Expect nil but got %v", actual)
	}
}

func Test_ValidateInMyTable_Y_4(t *testing.T) {
	position := robotwalk.Position{X: 0, Y: 4}
	actual := robotwalk.ValidateInMyTable(position)

	if actual != nil {
		t.Errorf("Expect nil but got %v", actual)
	}
}

func Test_ValidateInMyTable_Y_minus_4(t *testing.T) {
	position := robotwalk.Position{X: 0, Y: -4}
	actual := robotwalk.ValidateInMyTable(position)

	if actual != nil {
		t.Errorf("Expect nil but got %v", actual)
	}
}

func Test_ValidateInMyTable_X_4_Y_4(t *testing.T) {
	position := robotwalk.Position{X: -4, Y: 4}
	actual := robotwalk.ValidateInMyTable(position)

	if actual != nil {
		t.Errorf("Expect nil but got %v", actual)
	}
}
