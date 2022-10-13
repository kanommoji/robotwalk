package robotwalk_test

import (
	robotwalk "robotwalk/src/service"
	"testing"
)

func Test_InitWalking(t *testing.T) {
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

	robotwalk.InitWalking()
	actual := robotwalk.MyRobot

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_CalculateNextPosition_Walk_W_Direction_North_OutOfTable(t *testing.T) {
	walk := "W"
	robotwalk.MyRobot = robotwalk.Robot{
		Direction: robotwalk.North,
		Position:  robotwalk.Position{X: 0, Y: 4},
		Table:     robotwalk.Table{Row: 0, Column: 4},
	}
	expected := robotwalk.Robot{
		Direction: robotwalk.North,
		Position:  robotwalk.Position{X: 0, Y: 4},
		Table:     robotwalk.Table{Row: 0, Column: 4},
	}

	robotwalk.CalculateNextDirection(walk)
	actual := robotwalk.MyRobot

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_CalculateNextPosition_Walk_W_Direction_North(t *testing.T) {
	walk := "W"
	robotwalk.MyRobot = robotwalk.Robot{
		Direction: robotwalk.North,
		Position:  robotwalk.Position{X: 0, Y: 0},
		Table:     robotwalk.Table{Row: 4, Column: 4},
	}
	expected := robotwalk.Robot{
		Direction: robotwalk.North,
		Position:  robotwalk.Position{X: 0, Y: 1},
		Table:     robotwalk.Table{Row: 3, Column: 4},
	}

	robotwalk.CalculateNextDirection(walk)
	actual := robotwalk.MyRobot

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_CalculateNextPosition_Walk_W_Direction_East(t *testing.T) {
	walk := "W"
	robotwalk.MyRobot = robotwalk.Robot{
		Direction: robotwalk.East,
		Position:  robotwalk.Position{X: 0, Y: 0},
		Table:     robotwalk.Table{Row: 4, Column: 4},
	}
	expected := robotwalk.Robot{
		Direction: robotwalk.East,
		Position:  robotwalk.Position{X: 1, Y: 0},
		Table:     robotwalk.Table{Row: 4, Column: 5},
	}

	robotwalk.CalculateNextDirection(walk)
	actual := robotwalk.MyRobot

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_CalculateNextPosition_Walk_W_Direction_South(t *testing.T) {
	walk := "W"
	robotwalk.MyRobot = robotwalk.Robot{
		Direction: robotwalk.South,
		Position:  robotwalk.Position{X: 0, Y: 0},
		Table:     robotwalk.Table{Row: 4, Column: 4},
	}
	expected := robotwalk.Robot{
		Direction: robotwalk.South,
		Position:  robotwalk.Position{X: 0, Y: -1},
		Table:     robotwalk.Table{Row: 5, Column: 4},
	}

	robotwalk.CalculateNextDirection(walk)
	actual := robotwalk.MyRobot

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_CalculateNextPosition_Walk_W_Direction_West(t *testing.T) {
	walk := "W"
	robotwalk.MyRobot = robotwalk.Robot{
		Direction: robotwalk.West,
		Position:  robotwalk.Position{X: 0, Y: 0},
		Table:     robotwalk.Table{Row: 4, Column: 4},
	}
	expected := robotwalk.Robot{
		Direction: robotwalk.West,
		Position:  robotwalk.Position{X: -1, Y: 0},
		Table:     robotwalk.Table{Row: 4, Column: 3},
	}

	robotwalk.CalculateNextDirection(walk)
	actual := robotwalk.MyRobot

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_CalculateNextPosition_Walk_L_Change_Direction_North_To_West(t *testing.T) {
	walk := "L"
	robotwalk.MyRobot = robotwalk.Robot{Direction: robotwalk.North}
	expected := robotwalk.Robot{Direction: robotwalk.West}

	robotwalk.CalculateNextDirection(walk)
	actual := robotwalk.MyRobot

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_CalculateNextPosition_Walk_L_Change_Direction_West_To_South(t *testing.T) {
	walk := "L"
	robotwalk.MyRobot = robotwalk.Robot{Direction: robotwalk.West}
	expected := robotwalk.Robot{Direction: robotwalk.South}

	robotwalk.CalculateNextDirection(walk)
	actual := robotwalk.MyRobot

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_CalculateNextPosition_Walk_L_Change_Direction_South_North_To_East(t *testing.T) {
	walk := "L"
	robotwalk.MyRobot = robotwalk.Robot{Direction: robotwalk.South}
	expected := robotwalk.Robot{Direction: robotwalk.East}

	robotwalk.CalculateNextDirection(walk)
	actual := robotwalk.MyRobot

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_CalculateNextPosition_Walk_L_Change_Direction_East_To_North(t *testing.T) {
	walk := "L"
	robotwalk.MyRobot = robotwalk.Robot{Direction: robotwalk.East}
	expected := robotwalk.Robot{Direction: robotwalk.North}

	robotwalk.CalculateNextDirection(walk)
	actual := robotwalk.MyRobot

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_CalculateNextPosition_Walk_R_Change_Direction_North_To_East(t *testing.T) {
	walk := "R"
	robotwalk.MyRobot = robotwalk.Robot{Direction: robotwalk.North}
	expected := robotwalk.Robot{Direction: robotwalk.East}

	robotwalk.CalculateNextDirection(walk)
	actual := robotwalk.MyRobot

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_CalculateNextPosition_Walk_R_Change_Direction_East_To_South(t *testing.T) {
	walk := "R"
	robotwalk.MyRobot = robotwalk.Robot{Direction: robotwalk.East}
	expected := robotwalk.Robot{Direction: robotwalk.South}

	robotwalk.CalculateNextDirection(walk)
	actual := robotwalk.MyRobot

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_CalculateNextPosition_Walk_R_Change_Direction_South_To_West(t *testing.T) {
	walk := "R"
	robotwalk.MyRobot = robotwalk.Robot{Direction: robotwalk.South}
	expected := robotwalk.Robot{Direction: robotwalk.West}

	robotwalk.CalculateNextDirection(walk)
	actual := robotwalk.MyRobot

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_CalculateNextPosition_Walk_R_Change_Direction_West_To_North(t *testing.T) {
	walk := "R"
	robotwalk.MyRobot = robotwalk.Robot{Direction: robotwalk.West}
	expected := robotwalk.Robot{Direction: robotwalk.North}

	robotwalk.CalculateNextDirection(walk)
	actual := robotwalk.MyRobot

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}
