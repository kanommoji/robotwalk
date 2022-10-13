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

	actual, _ := robotwalk.CalculateNextDirection(walk, robot)

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

	actual, _ := robotwalk.CalculateNextDirection(walk, robot)

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

	actual, _ := robotwalk.CalculateNextDirection(walk, robot)

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

	actual, _ := robotwalk.CalculateNextDirection(walk, robot)

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

	actual, _ := robotwalk.CalculateNextDirection(walk, robot)

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_CalculateNextPosition_Walk_L_Change_Direction_North_To_West(t *testing.T) {
	walk := "L"
	robot := robotwalk.Robot{Direction: robotwalk.North}
	expected := robotwalk.Robot{Direction: robotwalk.West}

	actual, _ := robotwalk.CalculateNextDirection(walk, robot)

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_CalculateNextPosition_Walk_L_Change_Direction_West_To_South(t *testing.T) {
	walk := "L"
	robot := robotwalk.Robot{Direction: robotwalk.West}
	expected := robotwalk.Robot{Direction: robotwalk.South}

	actual, _ := robotwalk.CalculateNextDirection(walk, robot)

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_CalculateNextPosition_Walk_L_Change_Direction_South_North_To_East(t *testing.T) {
	walk := "L"
	robot := robotwalk.Robot{Direction: robotwalk.South}
	expected := robotwalk.Robot{Direction: robotwalk.East}

	actual, _ := robotwalk.CalculateNextDirection(walk, robot)

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_CalculateNextPosition_Walk_L_Change_Direction_East_To_North(t *testing.T) {
	walk := "L"
	robot := robotwalk.Robot{Direction: robotwalk.East}
	expected := robotwalk.Robot{Direction: robotwalk.North}

	actual, _ := robotwalk.CalculateNextDirection(walk, robot)

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_CalculateNextPosition_Walk_R_Change_Direction_North_To_East(t *testing.T) {
	walk := "R"
	robot := robotwalk.Robot{Direction: robotwalk.North}
	expected := robotwalk.Robot{Direction: robotwalk.East}

	actual, _ := robotwalk.CalculateNextDirection(walk, robot)

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_CalculateNextPosition_Walk_R_Change_Direction_East_To_South(t *testing.T) {
	walk := "R"
	robot := robotwalk.Robot{Direction: robotwalk.East}
	expected := robotwalk.Robot{Direction: robotwalk.South}

	actual, _ := robotwalk.CalculateNextDirection(walk, robot)

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_CalculateNextPosition_Walk_R_Change_Direction_South_To_West(t *testing.T) {
	walk := "R"
	robot := robotwalk.Robot{Direction: robotwalk.South}
	expected := robotwalk.Robot{Direction: robotwalk.West}

	actual, _ := robotwalk.CalculateNextDirection(walk, robot)

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_CalculateNextPosition_Walk_R_Change_Direction_West_To_North(t *testing.T) {
	walk := "R"
	robot := robotwalk.Robot{Direction: robotwalk.West}
	expected := robotwalk.Robot{Direction: robotwalk.North}

	actual, _ := robotwalk.CalculateNextDirection(walk, robot)

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}
